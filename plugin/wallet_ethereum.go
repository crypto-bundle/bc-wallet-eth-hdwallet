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
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/ethereum/go-ethereum/crypto"
)

// ethereumWallet parent
type ethereumWallet struct {
	purpose  int
	coinType int

	account       uint32
	change        uint32
	addressNumber uint32

	extendedKey *keyBundle
	accountKey  *accountKey

	blockChainParams *chaincfg.Params
}

// NewAccount create new account via mnemonic wallet by
func (w *wallet) NewAccount(account, change, index uint32) (*ethereumWallet, error) {
	accKey, extendedKey, err := w.GetChildKey(Bip44Purpose,
		uint32(pluginCoinType), account, change, index)
	if err != nil {
		return nil, err
	}

	return &ethereumWallet{
		purpose:          Bip44Purpose,
		coinType:         pluginCoinType,
		account:          account,
		change:           change,
		addressNumber:    index,
		extendedKey:      extendedKey,
		accountKey:       accKey,
		blockChainParams: w.Network,
	}, err
}

// GetAddress get address with 0x
func (e *ethereumWallet) GetAddress() (string, error) {
	return crypto.PubkeyToAddress(*e.extendedKey.PublicECDSA).Hex(), nil
}

// GetPubKey get key with 0x
func (e *ethereumWallet) GetPubKey() string {
	return "0x" + e.extendedKey.PublicHex()
}

// GetPrvKey get key with 0x
func (e *ethereumWallet) GetPrvKey() (string, error) {
	return "0x" + hex.EncodeToString(e.extendedKey.PrivateECDSA.D.Bytes()), nil
}

// GetPath ...
func (e *ethereumWallet) GetPath() string {
	return fmt.Sprintf("m/%d'/%d'/%d'/%d/%d",
		e.GetPurpose(), e.GetCoinType(), e.account, e.change, e.addressNumber)
}

// GetPurpose ...
func (e *ethereumWallet) GetPurpose() int {
	return e.purpose
}

// GetCoinType ...
func (e *ethereumWallet) GetCoinType() int {
	return ethereumCoinNumber
}

func (e *ethereumWallet) CloneECDSAPrivateKey() *ecdsa.PrivateKey {
	return e.extendedKey.CloneECDSAPrivateKey()
}

func (e *ethereumWallet) ClearSecrets() {
	e.accountKey.ExtendedKey.Zero()
	e.accountKey.Public = ""
	e.accountKey.Private = ""

	e.extendedKey.ClearSecrets()

	e.blockChainParams = nil
	e.accountKey.Network = nil
	e.accountKey = nil
	e.extendedKey = nil
	e.account = 0
	e.change = 0
	e.addressNumber = 0
}
