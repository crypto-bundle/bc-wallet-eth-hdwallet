/*
 * MIT License
 *
 * Copyright (c) 2021-2023 Aleksei Kotelnikov
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package main

import (
	"errors"
	"fmt"
	"math/big"
	"sync"

	pbEthereum "github.com/crypto-bundle/bc-wallet-eth-hdwallet/pkg/proto"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

var (
	ErrMissingValue        = errors.New("missing required values in tx data")
	ErrUnsupportedDataType = errors.New("unsupported signature data type")
)

var onceMarshaller sync.Once
var marshallerSvc *signDataMarshaller

type signDataMarshaller struct {
}

func (m *signDataMarshaller) MarshallSignData(dataForSign []byte) (types.TxData, error) {
	pbAny := &anypb.Any{}
	err := proto.Unmarshal(dataForSign, pbAny)
	if err != nil {
		return nil, err
	}

	switch pbAny.TypeUrl {
	case "ethereum.LegacyTxData":
		return m.marshallEthereumLegacyTx(pbAny)
	case "ethereum.DynamicFeeTxData":
		return nil, ErrUnsupportedDataType
		//return m.marshallEthereumDynamicFeeTx(pbAny)
	case "ethereum.AccessListTxData":
		return nil, ErrUnsupportedDataType
		//return m.marshallEthereumAccessListTx(pbAny)
	default:
		return nil, ErrUnsupportedDataType
	}
}

func (m *signDataMarshaller) marshallEthereumLegacyTx(dataForSign *anypb.Any) (types.TxData, error) {
	pbTx := &pbEthereum.LegacyTxData{}
	err := dataForSign.UnmarshalTo(pbTx)
	if err != nil {
		return nil, err
	}

	if pbTx.ToAddress == nil {
		return nil, fmt.Errorf("%w: %s", ErrMissingValue, "destination address - To")
	}

	if pbTx.GasPrice == nil {
		return nil, fmt.Errorf("%w: %s", ErrMissingValue, "gas price - GasPrice")
	}

	if pbTx.Value == nil {
		return nil, fmt.Errorf("%w: %s", ErrMissingValue, "transfer amount - Value")
	}

	ethAddr := common.BytesToAddress(pbTx.ToAddress)

	tx := &types.LegacyTx{
		Nonce:    pbTx.Nonce,
		GasPrice: big.NewInt(0).SetBytes(pbTx.GasPrice),
		Gas:      pbTx.GasLimit,
		To:       &ethAddr,
		Value:    big.NewInt(0).SetBytes(pbTx.Value),
		Data:     pbTx.Data,
		V:        nil,
		R:        nil,
		S:        nil,
	}

	if pbTx.SignParameters != nil {
		if pbTx.SignParameters.V != nil {
			tx.V = big.NewInt(0).SetBytes(pbTx.SignParameters.V)
		}

		if pbTx.SignParameters.R != nil {
			tx.R = big.NewInt(0).SetBytes(pbTx.SignParameters.R)
		}

		if pbTx.SignParameters.S != nil {
			tx.S = big.NewInt(0).SetBytes(pbTx.SignParameters.S)
		}
	}

	return tx, nil
}

func (m *signDataMarshaller) marshallEthereumDynamicFeeTx(dataForSign *anypb.Any) (types.TxData, error) {
	pbTx := &pbEthereum.DynamicFeeTxData{}
	err := dataForSign.UnmarshalTo(pbTx)
	if err != nil {
		return nil, err
	}

	if pbTx.ToAddress == nil {
		return nil, fmt.Errorf("%w: %s", ErrMissingValue, "destination address - To")
	}

	if pbTx.GasFeeCap == nil {
		return nil, fmt.Errorf("%w: %s", ErrMissingValue, "gas fee capacity - GasFeeCap")
	}

	if pbTx.GasTipCap == nil {
		return nil, fmt.Errorf("%w: %s", ErrMissingValue, "gas tips - GasTipCap")
	}

	if pbTx.Value == nil {
		return nil, fmt.Errorf("%w: %s", ErrMissingValue, "transfer amount - Value")
	}

	ethAddr := common.BytesToAddress(pbTx.ToAddress)

	tx := &types.DynamicFeeTx{
		Nonce:      pbTx.Nonce,
		GasFeeCap:  big.NewInt(0).SetBytes(pbTx.GasFeeCap),
		GasTipCap:  big.NewInt(0).SetBytes(pbTx.GasTipCap),
		Gas:        pbTx.GasLimit,
		To:         &ethAddr,
		Value:      big.NewInt(0).SetBytes(pbTx.Value),
		Data:       pbTx.Data,
		AccessList: nil,
		V:          nil,
		R:          nil,
		S:          nil,
	}

	if pbTx.AccessList != nil {
		acList := make(types.AccessList, len(pbTx.AccessList))

		for i, _ := range pbTx.AccessList {
			pbTuple := pbTx.AccessList[i]

			tupleAddr := common.BytesToAddress(pbTuple.Address)
			accessListTuple := types.AccessTuple{
				Address:     tupleAddr,
				StorageKeys: make([]common.Hash, len(pbTuple.StorageKeys)),
			}

			for j, _ := range pbTuple.StorageKeys {
				storageKey := pbTuple.StorageKeys[j]

				accessListTuple.StorageKeys[j] = common.BytesToHash(storageKey)
			}

			acList[i] = accessListTuple
		}

		tx.AccessList = acList
	}

	if pbTx.SignParameters != nil {
		if pbTx.SignParameters.V != nil {
			tx.V = big.NewInt(0).SetBytes(pbTx.SignParameters.V)
		}

		if pbTx.SignParameters.R != nil {
			tx.R = big.NewInt(0).SetBytes(pbTx.SignParameters.R)
		}

		if pbTx.SignParameters.S != nil {
			tx.S = big.NewInt(0).SetBytes(pbTx.SignParameters.S)
		}
	}

	return tx, nil
}

func (m *signDataMarshaller) marshallEthereumAccessListTx(dataForSign *anypb.Any) (types.TxData, error) {
	pbTx := &pbEthereum.AccessListTxData{}
	err := dataForSign.UnmarshalTo(pbTx)
	if err != nil {
		return nil, err
	}

	if pbTx.ToAddress == nil {
		return nil, fmt.Errorf("%w: %s", ErrMissingValue, "destination address - To")
	}

	if pbTx.GasPrice == nil {
		return nil, fmt.Errorf("%w: %s", ErrMissingValue, "gas price - GasPrice")
	}

	if pbTx.Value == nil {
		return nil, fmt.Errorf("%w: %s", ErrMissingValue, "transfer amount - Value")
	}

	accessListTuplesCount := len(pbTx.AccessList)
	if accessListTuplesCount == 0 {
		return nil, fmt.Errorf("%w: %s", ErrMissingValue, "access list is empty - AccessList")
	}

	ethAddr := common.BytesToAddress(pbTx.ToAddress)

	tx := &types.AccessListTx{
		Nonce:      pbTx.Nonce,
		GasPrice:   big.NewInt(0).SetBytes(pbTx.GasPrice),
		Gas:        pbTx.GasLimit,
		To:         &ethAddr,
		Value:      big.NewInt(0).SetBytes(pbTx.Value),
		Data:       pbTx.Data,
		AccessList: make(types.AccessList, accessListTuplesCount),
		V:          nil,
		R:          nil,
		S:          nil,
	}

	for i, _ := range pbTx.AccessList {
		pbTuple := pbTx.AccessList[i]

		tupleAddr := common.BytesToAddress(pbTuple.Address)
		accessListTuple := types.AccessTuple{
			Address:     tupleAddr,
			StorageKeys: make([]common.Hash, len(pbTuple.StorageKeys)),
		}

		for j, _ := range pbTuple.StorageKeys {
			storageKey := pbTuple.StorageKeys[j]

			accessListTuple.StorageKeys[j] = common.BytesToHash(storageKey)
		}

		tx.AccessList[i] = accessListTuple
	}

	if pbTx.SignParameters != nil {
		if pbTx.SignParameters.V != nil {
			tx.V = big.NewInt(0).SetBytes(pbTx.SignParameters.V)
		}

		if pbTx.SignParameters.R != nil {
			tx.R = big.NewInt(0).SetBytes(pbTx.SignParameters.R)
		}

		if pbTx.SignParameters.S != nil {
			tx.S = big.NewInt(0).SetBytes(pbTx.SignParameters.S)
		}
	}

	return tx, nil
}

func newMarshallerService() *signDataMarshaller {
	onceMarshaller.Do(func() {
		marshallerSvc = &signDataMarshaller{}
	})

	return marshallerSvc
}
