/*
 *
 *
 * MIT NON-AI License
 *
 * Copyright (c) 2022-2024 Aleksei Kotelnikov(gudron2s@gmail.com)
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of the software and associated documentation files (the "Software"),
 * to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense,
 * and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions.
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 *
 * In addition, the following restrictions apply:
 *
 * 1. The Software and any modifications made to it may not be used for the purpose of training or improving machine learning algorithms,
 * including but not limited to artificial intelligence, natural language processing, or data mining. This condition applies to any derivatives,
 * modifications, or updates based on the Software code. Any usage of the Software in an AI-training dataset is considered a breach of this License.
 *
 * 2. The Software may not be included in any dataset used for training or improving machine learning algorithms,
 * including but not limited to artificial intelligence, natural language processing, or data mining.
 *
 * 3. Any person or organization found to be in violation of these restrictions will be subject to legal action and may be held liable
 * for any damages resulting from such use.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
 * DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
 * OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 *
 */

package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"sync"

	pbCommon "github.com/crypto-bundle/bc-wallet-common-hdwallet-controller/pkg/grpc/common"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/ethereum/go-ethereum/core/types"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

const addrPatKeyTemplate = "%d'/%d/%d"

type addressData struct {
	address    string
	privateKey *ecdsa.PrivateKey
}

func (e *addressData) ClonePrivateKey() *ecdsa.PrivateKey {
	return &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: btcec.S256(),
			X:     (&big.Int{}).SetBytes(e.privateKey.X.Bytes()),
			Y:     (&big.Int{}).SetBytes(e.privateKey.Y.Bytes()),
		},
		D: (&big.Int{}).SetBytes(e.privateKey.D.Bytes()),
	}
}

type mnemonicWalletUnit struct {
	mu *sync.Mutex

	hdWalletSvc *wallet
	dataSigner  types.Signer

	mnemonicWalletUUID string
	mnemonicHash       string

	// addressPool is pool of derivation addresses with private keys and address
	// map key - string with derivation path
	// map value - ecdsa.PrivateKey and address string
	addressPool map[string]*addressData
}

func (u *mnemonicWalletUnit) Shutdown(ctx context.Context) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	err := u.unloadWallet()
	if err != nil {
		return fmt.Errorf("unable to unload wallet: %w", err)
	}

	return nil
}

func (u *mnemonicWalletUnit) UnloadWallet() error {
	u.mu.Lock()
	defer u.mu.Unlock()

	return u.unloadWallet()
}

func (u *mnemonicWalletUnit) unloadWallet() error {
	for accountPath, _ := range u.addressPool {
		addrData, isExist := u.addressPool[accountPath]
		if !isExist {
			continue
		}

		if addrData == nil {
			continue
		}

		if addrData.privateKey != nil {
			zeroKey(addrData.privateKey)
		}
		addrData.address = ""

		u.addressPool[accountPath] = nil

		delete(u.addressPool, accountPath)
	}

	u.addressPool = nil
	u.mnemonicWalletUUID = "0"
	u.mnemonicHash = "0"

	u.hdWalletSvc.ClearSecrets()
	u.hdWalletSvc = nil

	return nil
}

func (u *mnemonicWalletUnit) GetWalletUUID() string {
	return u.mnemonicWalletUUID
}

func (u *mnemonicWalletUnit) SignData(ctx context.Context,
	accountParameters *anypb.Any,
	dataForSign []byte,
) (*string, []byte, error) {
	accIdentity := &pbCommon.DerivationAddressIdentity{}
	err := proto.Unmarshal(accountParameters.GetValue(), accIdentity)
	if err != nil {
		return nil, nil, err
	}

	tx := &types.Transaction{}
	err = tx.UnmarshalBinary(dataForSign)
	if err != nil {
		return nil, nil, err
	}

	u.mu.Lock()
	defer u.mu.Unlock()

	return u.signData(ctx,
		accIdentity.AccountIndex,
		accIdentity.InternalIndex,
		accIdentity.AddressIndex,
		tx)
}

func (u *mnemonicWalletUnit) signData(ctx context.Context,
	account, change, index uint32,
	txForSign *types.Transaction,
) (*string, []byte, error) {
	addr, privKey, err := u.loadAccountDataByPath(ctx, account, change, index)
	if err != nil {
		return nil, nil, err
	}

	signedTx, err := types.SignTx(txForSign, u.dataSigner, privKey)
	if err != nil {
		return nil, nil, err
	}

	signedTxRawData, err := signedTx.MarshalBinary()
	if err != nil {
		return nil, nil, fmt.Errorf("unable to sign: %w", err)
	}

	return addr, signedTxRawData, nil
}

func (u *mnemonicWalletUnit) LoadAccount(ctx context.Context,
	accountParameters *anypb.Any,
) (*string, error) {
	accIdentity := &pbCommon.DerivationAddressIdentity{}
	err := accountParameters.UnmarshalTo(accIdentity)
	if err != nil {
		return nil, err
	}

	u.mu.Lock()
	defer u.mu.Unlock()

	addrData, _, err := u.loadAccountDataByPath(ctx, accIdentity.AccountIndex,
		accIdentity.InternalIndex,
		accIdentity.AddressIndex)
	if err != nil {
		return nil, err
	}

	if addrData == nil {
		return nil, nil
	}

	return addrData, nil
}

func (u *mnemonicWalletUnit) loadAccountDataByPath(ctx context.Context,
	account, change, index uint32,
) (*string, *ecdsa.PrivateKey, error) {
	mapKey := fmt.Sprintf(addrPatKeyTemplate, account, change, index)
	addrData, isExists := u.addressPool[mapKey]
	if !isExists {
		hdWalletAccount, walletErr := u.hdWalletSvc.NewAccount(account, change, index)
		if walletErr != nil {
			return nil, nil, walletErr
		}

		addr, walletErr := hdWalletAccount.GetAddress()
		if walletErr != nil {
			return nil, nil, walletErr
		}

		addrData = &addressData{
			address:    addr,
			privateKey: hdWalletAccount.CloneECDSAPrivateKey(),
		}
		u.addressPool[mapKey] = addrData

		hdWalletAccount.ClearSecrets()
		hdWalletAccount = nil
	}

	return &addrData.address, addrData.ClonePrivateKey(), nil
}

func (u *mnemonicWalletUnit) GetAccountAddress(ctx context.Context,
	accountParameters *anypb.Any,
) (*string, error) {
	accIdentity := &pbCommon.DerivationAddressIdentity{}
	err := accountParameters.UnmarshalTo(accIdentity)
	if err != nil {
		return nil, err
	}

	return u.getAddressByPath(ctx, accIdentity.AccountIndex,
		accIdentity.InternalIndex,
		accIdentity.AddressIndex)
}

func (u *mnemonicWalletUnit) GetMultipleAccounts(ctx context.Context,
	multipleAccountsParameters *anypb.Any,
) (uint, []*pbCommon.AccountIdentity, error) {
	list := &pbCommon.RangeUnitsList{}
	err := multipleAccountsParameters.UnmarshalTo(list)
	if err != nil {
		return 0, nil, err
	}

	return u.getMultipleAccounts(ctx, list)
}

func (u *mnemonicWalletUnit) getMultipleAccounts(ctx context.Context,
	rangeList *pbCommon.RangeUnitsList,
) (uint, []*pbCommon.AccountIdentity, error) {
	var err error
	size := len(rangeList.RangeUnits)

	result := make([]*pbCommon.AccountIdentity, 0)
	var resCount uint

	for i := uint32(0); i != uint32(size); i++ {
		rangeUnit := rangeList.RangeUnits[i]

		count, list, loopErr := u.getAccountsByRange(ctx, rangeUnit)
		if loopErr != nil {
			return 0, nil, loopErr
		}

		resCount += count
		result = append(result, list...)
	}

	if err != nil {
		return 0, nil, err
	}

	return resCount, result, nil
}

func (u *mnemonicWalletUnit) getAccountsByRange(ctx context.Context,
	rangeUnit *pbCommon.RangeRequestUnit,
) (uint, []*pbCommon.AccountIdentity, error) {
	diff := rangeUnit.AddressIndexTo - rangeUnit.AddressIndexFrom
	if diff == 0 { // if one item in range
		accountIdentifier, loopErr := u.getAddressAndMarshal(ctx, rangeUnit.AccountIndex,
			rangeUnit.InternalIndex, rangeUnit.AddressIndexFrom)
		if loopErr != nil {
			return 0, nil, loopErr
		}

		return 1, []*pbCommon.AccountIdentity{accountIdentifier}, nil
	}
	elementsCount := diff + 1
	result := make([]*pbCommon.AccountIdentity, elementsCount)

	wg := sync.WaitGroup{}
	wg.Add(int(elementsCount))
	var position uint32

	for addressIndex := rangeUnit.AddressIndexFrom; addressIndex <= rangeUnit.AddressIndexTo; addressIndex++ {
		go func(accountIdx, internalIdx, addressIdx, position uint32) {
			defer wg.Done()

			accountIdentifier, loopErr := u.getAddressAndMarshal(ctx, rangeUnit.AccountIndex,
				rangeUnit.InternalIndex, addressIdx)
			if loopErr != nil {
				return
			}

			result[position] = accountIdentifier

			return
		}(rangeUnit.AccountIndex, rangeUnit.InternalIndex, addressIndex, position)

		position++
	}

	wg.Wait()

	return uint(len(result)), result, nil
}

func (u *mnemonicWalletUnit) getAddressAndMarshal(ctx context.Context,
	account, change, index uint32,
) (*pbCommon.AccountIdentity, error) {
	address, err := u.getAddressByPath(ctx, account,
		change, index)
	if err != nil {
		return nil, err
	}

	addrParams := &anypb.Any{}
	err = addrParams.MarshalFrom(&pbCommon.DerivationAddressIdentity{
		AccountIndex:  account,
		InternalIndex: change,
		AddressIndex:  index,
	})
	if err != nil {
		return nil, err
	}

	return &pbCommon.AccountIdentity{
		Parameters: addrParams,
		Address:    *address,
	}, nil
}

func (u *mnemonicWalletUnit) getAddressByPath(_ context.Context,
	account, change, index uint32,
) (*string, error) {
	hdWalletAccount, err := u.hdWalletSvc.NewAccount(account, change, index)
	if err != nil {
		return nil, err
	}

	defer func() {
		hdWalletAccount.ClearSecrets()
		hdWalletAccount = nil
	}()

	blockchainAddress, err := hdWalletAccount.GetAddress()
	if err != nil {
		return nil, err
	}

	return &blockchainAddress, nil
}

func NewPoolUnit(walletUUID string,
	mnemonicDecryptedData string,
) (interface{}, error) {
	hdWalletSvc, createErr := newWalletFromMnemonic(mnemonicDecryptedData)
	if createErr != nil {
		return nil, createErr
	}

	return &mnemonicWalletUnit{
		mu: &sync.Mutex{},

		hdWalletSvc: hdWalletSvc,

		dataSigner: pluginSigner,

		mnemonicWalletUUID: walletUUID,

		addressPool: make(map[string]*addressData),
	}, nil
}
