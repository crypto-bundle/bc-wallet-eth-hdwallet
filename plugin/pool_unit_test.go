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
	"bytes"
	"context"
	"fmt"
	"log"

	"math/big"
	"testing"

	pbCommon "github.com/crypto-bundle/bc-wallet-common-hdwallet-controller/pkg/grpc/common"
	pbEthereum "github.com/crypto-bundle/bc-wallet-eth-hdwallet/pkg/proto"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/tyler-smith/go-bip39"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func TestNewPoolUnit(t *testing.T) {
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		t.Fatalf("%s: %e", "unable to create entropy:", err)
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		t.Fatalf("%s: %e", "unable to create mnemonic pharase from entory:", err)
	}

	_, err = NewPoolUnit(uuid.NewString(), mnemonic)
	if err != nil {
		t.Fatalf("%s: %e", "unable to create mnemonic wallet pool unit:", err)
	}
}

func TestMnemonicWalletUnit_GetWalletUUID(t *testing.T) {
	type testCase struct {
		WalletUUID  string
		Mnemonic    string
		AddressPath *pbCommon.DerivationAddressIdentity

		ExpectedAddress string
	}

	// WARN: DO NOT USE THIS MNEMONIC IN MAINNET OR TESTNET. Usage only in unit-tests
	// WARN: DO NOT USE THIS MNEMONIC IN MAINNET OR TESTNET. Usage only in unit-tests
	// WARN: DO NOT USE THIS MNEMONIC IN MAINNET OR TESTNET. Usage only in unit-tests
	tCase := &testCase{
		WalletUUID: uuid.NewString(),
		Mnemonic:   "seven kitten wire trap family giraffe globe access dinosaur upper forum aerobic dash segment cruise concert giant upon sniff armed rain royal firm state",
		AddressPath: &pbCommon.DerivationAddressIdentity{
			AccountIndex:  5,
			InternalIndex: 5,
			AddressIndex:  55,
		},
		ExpectedAddress: "0xec9bddc83CC1F391Db7e7853FF7E444B07c0691d",
	}

	poolUnitIntrf, loopErr := NewPoolUnit(tCase.WalletUUID, tCase.Mnemonic)
	if loopErr != nil {
		t.Fatalf("%s: %e", "unable to create mnemonic wallet pool unit:", loopErr)
	}

	poolUnit, ok := poolUnitIntrf.(*mnemonicWalletUnit)
	if !ok {
		t.Fatalf("%s", "unable to cast interface to pool unit worker")
	}

	accountIdentity := &anypb.Any{}
	_ = accountIdentity.MarshalFrom(tCase.AddressPath)

	addr, loopErr := poolUnit.GetAccountAddress(context.Background(), accountIdentity)
	if loopErr != nil {
		t.Fatalf("%s: %e", "unable to get address from pool unit", loopErr)
	}

	if addr == nil {
		t.Fatalf("%s", "missing address in pool unit result")
	}

	if tCase.ExpectedAddress != *addr {
		t.Fatalf("%s", "address not equal with expected")
	}

	resultUUID := poolUnit.GetWalletUUID()
	if tCase.WalletUUID != resultUUID {
		t.Fatalf("%s", "wallet uuid not equal with expected")
	}
}

func TestMnemonicWalletUnit_GetAccountAddress(t *testing.T) {
	type testCase struct {
		Mnemonic    string
		AddressPath *pbCommon.DerivationAddressIdentity

		ExpectedAddress string
	}

	// WARN: DO NOT USE THESE MNEMONICS IN MAINNET OR TESTNET. Usage only in unit-tests
	// WARN: DO NOT USE THESE MNEMONICS IN MAINNET OR TESTNET. Usage only in unit-tests
	// WARN: DO NOT USE THESE MNEMONICS IN MAINNET OR TESTNET. Usage only in unit-tests
	testCases := []*testCase{
		{
			Mnemonic: "unfair silver dune air rib enforce protect limit jazz dinner thumb drift spring warrior bonus snack argue flavor wild faculty derive open dynamic carpet",
			AddressPath: &pbCommon.DerivationAddressIdentity{
				AccountIndex:  3,
				InternalIndex: 13,
				AddressIndex:  114,
			},
			ExpectedAddress: "0xb96b5c70ff4102d0900E9Fc0614E5BA4FE486281",
		},
		{
			Mnemonic: "obscure town quick bundle north message want sketch brass tone vast spoil home gentle field ozone mushroom current math cat canvas plunge stay truly",
			AddressPath: &pbCommon.DerivationAddressIdentity{
				AccountIndex:  1020,
				InternalIndex: 10300,
				AddressIndex:  104000,
			},
			ExpectedAddress: "0xC342507376b862F9f09845Db10915c256473Bd9e",
		},
		{
			Mnemonic: "beach large spray gentle buyer hover flock dream hybrid match whip ten mountain pitch enemy lobster afford barrel patrol desk trigger output excuse truck",
			AddressPath: &pbCommon.DerivationAddressIdentity{
				AccountIndex:  2,
				InternalIndex: 104,
				AddressIndex:  1005,
			},
			ExpectedAddress: "0x7CD20e3Cf394F1C53316782BA553F2A37Ac93618",
		},
	}

	for _, tCase := range testCases {
		poolUnitIntrf, loopErr := NewPoolUnit(uuid.NewString(), tCase.Mnemonic)
		if loopErr != nil {
			t.Fatalf("%s: %e", "unable to create mnemonic wallet pool unit:", loopErr)
		}

		poolUnit, ok := poolUnitIntrf.(*mnemonicWalletUnit)
		if !ok {
			t.Fatalf("%s", "unable to cast interface to pool unit worker")
		}

		accountIdentity := &anypb.Any{}
		_ = accountIdentity.MarshalFrom(tCase.AddressPath)

		addr, loopErr := poolUnit.GetAccountAddress(context.Background(), accountIdentity)
		if loopErr != nil {
			t.Fatalf("%s: %e", "unable to get address from pool unit:", loopErr)
		}

		if addr == nil {
			t.Fatalf("%s", "missing address in pool unit result")
		}

		if tCase.ExpectedAddress != *addr {
			t.Fatalf("%s", "address not equal with expected")
		}

		loopErr = poolUnit.Shutdown(context.Background())
		if loopErr != nil {
			t.Fatalf("%s", "unable to shurdown pool unit")
		}
	}
}

func TestMnemonicWalletUnit_GetMultipleAccounts_3_by_50(t *testing.T) {
	type testCase struct {
		Mnemonic        string
		AddressPathList *pbCommon.RangeUnitsList
	}

	// WARN: DO NOT USE THESE MNEMONICS IN MAINNET OR TESTNET. Usage only in unit-tests
	// WARN: DO NOT USE THESE MNEMONICS IN MAINNET OR TESTNET. Usage only in unit-tests
	// WARN: DO NOT USE THESE MNEMONICS IN MAINNET OR TESTNET. Usage only in unit-tests
	testCases := []*testCase{
		{
			Mnemonic: "web account soft juice relief green account rebel rifle gun follow thunder ski credit judge off educate round advice allow wink bitter first color",
			AddressPathList: &pbCommon.RangeUnitsList{
				RangeUnits: []*pbCommon.RangeRequestUnit{
					{AccountIndex: 3, InternalIndex: 8, AddressIndexFrom: 150, AddressIndexTo: 200},
					{AccountIndex: 155, InternalIndex: 5, AddressIndexFrom: 4, AddressIndexTo: 54},
					{AccountIndex: 2555, InternalIndex: 50, AddressIndexFrom: 250, AddressIndexTo: 300},
				},
			},
		},
	}

	for _, tCase := range testCases {
		poolUnitIntrf, loopErr := NewPoolUnit(uuid.NewString(), tCase.Mnemonic)
		if loopErr != nil {
			t.Fatalf("%s: %e", "unable to create mnemonic wallet pool unit:", loopErr)
		}

		poolUnit, ok := poolUnitIntrf.(*mnemonicWalletUnit)
		if !ok {
			t.Fatalf("%s", "unable to cast interface to pool unit worker")
		}

		anyRangeUnit := &anypb.Any{}
		err := anyRangeUnit.MarshalFrom(tCase.AddressPathList)
		if err != nil {
			t.Fatalf("%s", "unable to marshal request units list")
		}

		count, addrList, loopErr := poolUnit.GetMultipleAccounts(context.Background(), anyRangeUnit)
		if loopErr != nil {
			t.Fatalf("%s: %e", "unable to get address from pool unit:", loopErr)
		}

		if count == 0 {
			t.Fatalf("%s", "addr list count not equal with expected")
		}

		if addrList == nil {
			t.Fatalf("%s", "missing addr lsit in pool unit result")
		}

		if count != uint(len(addrList)) {
			t.Fatalf("%s", "length of addrlist count not equal with count rseult value")
		}

		if count != 153 {
			t.Fatalf("%s", "count value or getAccount result not equal with values count of expected map")
		}

		err = poolUnit.Shutdown(context.Background())
		if err != nil {
			t.Fatalf("%s", "unable to shurdown pool unit")
		}
	}
}

func TestMnemonicWalletUnit_GetMultipleAccounts(t *testing.T) {
	type testCase struct {
		Mnemonic        string
		AddressPathList *pbCommon.RangeUnitsList

		ExpectedAddress map[string]*pbCommon.DerivationAddressIdentity
	}

	// WARN: DO NOT USE THESE MNEMONICS IN MAINNET OR TESTNET. Usage only in unit-tests
	// WARN: DO NOT USE THESE MNEMONICS IN MAINNET OR TESTNET. Usage only in unit-tests
	// WARN: DO NOT USE THESE MNEMONICS IN MAINNET OR TESTNET. Usage only in unit-tests
	testCases := []*testCase{
		{
			Mnemonic: "web account soft juice relief green account rebel rifle gun follow thunder ski credit judge off educate round advice allow wink bitter first color",
			AddressPathList: &pbCommon.RangeUnitsList{
				RangeUnits: []*pbCommon.RangeRequestUnit{
					{AccountIndex: 3, InternalIndex: 8, AddressIndexFrom: 114, AddressIndexTo: 124},
					{AccountIndex: 155, InternalIndex: 5, AddressIndexFrom: 4, AddressIndexTo: 8},
					{AccountIndex: 2555, InternalIndex: 50, AddressIndexFrom: 250, AddressIndexTo: 255},
				},
			},
			ExpectedAddress: map[string]*pbCommon.DerivationAddressIdentity{
				"0x3c601d8b930Be1935495747A278b63fBAA26d905": {AccountIndex: 3, InternalIndex: 8, AddressIndex: 114},
				"0xA3e51Ee437F0A0150E5eB25eFe353d45C4162cb2": {AccountIndex: 3, InternalIndex: 8, AddressIndex: 115},
				"0x3D88dF24e098108C8147Fe25856B064b30c92A99": {AccountIndex: 3, InternalIndex: 8, AddressIndex: 116},
				"0xBD3B46A13A8a3d312989F700F7C2781dEe2D3CEe": {AccountIndex: 3, InternalIndex: 8, AddressIndex: 117},
				"0xc737244E70A749B218E8AF0771c862bF10A724b4": {AccountIndex: 3, InternalIndex: 8, AddressIndex: 118},
				"0x40477746fd12171626DADB9fD4E0fd6a6CB12bfA": {AccountIndex: 3, InternalIndex: 8, AddressIndex: 119},
				"0x2d2f4760b6D7A98e7E79E2322d07BF10304537Bb": {AccountIndex: 3, InternalIndex: 8, AddressIndex: 120},
				"0x082e99DE321005e16e08b09cd28FEEfA2609e4D7": {AccountIndex: 3, InternalIndex: 8, AddressIndex: 121},
				"0x6e8C67C41CE8B3b0Fd5cCC0b8728a0E4a22Cae29": {AccountIndex: 3, InternalIndex: 8, AddressIndex: 122},
				"0xE025BC9e69F040247375f177ed7633545C52E1be": {AccountIndex: 3, InternalIndex: 8, AddressIndex: 123},
				"0xb3D6C84FBd2d5E71d7EA86d99dbd45bAa2A94Ea9": {AccountIndex: 3, InternalIndex: 8, AddressIndex: 124},
				//
				"0x0EEDf8b3BFdef8C13f2353b7059B9f5F8CA80871": {AccountIndex: 155, InternalIndex: 5, AddressIndex: 4},
				"0xB7D7110da3e46b44a44138dEE8ee52d086C12f15": {AccountIndex: 155, InternalIndex: 5, AddressIndex: 5},
				"0x6790741b690c988765D583ec744A57C2e043e54e": {AccountIndex: 155, InternalIndex: 5, AddressIndex: 6},
				"0x794434525BcCA3ba1E1393Cc89756fE854e64317": {AccountIndex: 155, InternalIndex: 5, AddressIndex: 7},
				"0x8BB042AE94FeeD0cCC24C03A6d726Fa92cbfC23B": {AccountIndex: 155, InternalIndex: 5, AddressIndex: 8},
				//
				"0x5B39D2679ae8Ea338B5fe0Cd95502DaA804af268": {AccountIndex: 2555,
					InternalIndex: 50,
					AddressIndex:  250},
				"0x5B7a0d6C846001F611498Df95b2Db0614E89ab23": {AccountIndex: 2555,
					InternalIndex: 50,
					AddressIndex:  251},
				"0xc8808DCd66962e8fbbd0bd1B8dB2D8aCA8284cD9": {AccountIndex: 2555,
					InternalIndex: 50,
					AddressIndex:  252},
				"0xEF98034363821a9Dd6f9619A58E56CB7D208A9E5": {AccountIndex: 2555,
					InternalIndex: 50,
					AddressIndex:  253},
				"0x4C222F221E0e0009edAD1eaB05ea4574f490B2dB": {AccountIndex: 2555,
					InternalIndex: 50,
					AddressIndex:  254},
				"0xA0F6637C3083a07Ae16162D720d6dd65070B8Cc4": {AccountIndex: 2555,
					InternalIndex: 50,
					AddressIndex:  255},
			},
		},
	}

	for _, tCase := range testCases {
		poolUnitIntrf, loopErr := NewPoolUnit(uuid.NewString(), tCase.Mnemonic)
		if loopErr != nil {
			t.Fatalf("%s: %e", "unable to create mnemonic wallet pool unit:", loopErr)
		}

		poolUnit, ok := poolUnitIntrf.(*mnemonicWalletUnit)
		if !ok {
			t.Fatalf("%s", "unable to cast interface to pool unit worker")
		}

		anyRangeUnit := &anypb.Any{}
		err := anyRangeUnit.MarshalFrom(tCase.AddressPathList)
		if err != nil {
			t.Fatalf("%s", "unable to marshal request units list")
		}

		count, addrList, loopErr := poolUnit.GetMultipleAccounts(context.Background(), anyRangeUnit)
		if loopErr != nil {
			t.Fatalf("%s: %e", "unable to get address from pool unit:", loopErr)
		}

		if count == 0 {
			t.Fatalf("%s", "addr list count not equal with expected")
		}

		if addrList == nil {
			t.Fatalf("%s", "missing addr lsit in pool unit result")
		}

		if count != uint(len(addrList)) {
			t.Fatalf("%s", "length of addrlist count not equal with count rseult value")
		}

		if count != uint(len(tCase.ExpectedAddress)) {
			t.Fatalf("%s", "count value or getAccount result not equal with values count of expected map")
		}

		for i := 0; i != len(addrList); i++ {
			addr := addrList[i]

			accIdentifier, isExists := tCase.ExpectedAddress[addr.Address]
			if !isExists {
				t.Fatalf("%s", "missing addr in expected map")
			}

			marshaledAccIdentifier := &pbCommon.DerivationAddressIdentity{}
			marshalErr := addr.Parameters.UnmarshalTo(marshaledAccIdentifier)
			if marshalErr != nil {
				t.Fatalf("%s: %e", "unable to unmarshal anydata to account identity", marshalErr)
			}

			if accIdentifier.AddressIndex != marshaledAccIdentifier.AddressIndex {
				t.Fatalf("%s", "marshaled address index not equal with expected")
			}

			if accIdentifier.InternalIndex != marshaledAccIdentifier.InternalIndex {
				t.Fatalf("%s", "marshaled internal index not equal with expected")
			}

			if accIdentifier.AccountIndex != marshaledAccIdentifier.AccountIndex {
				t.Fatalf("%s", "marshaled account index not equal with expected")
			}
		}

		err = poolUnit.Shutdown(context.Background())
		if err != nil {
			t.Fatalf("%s", "unable to shurdown pool unit")
		}
	}
}

func TestMnemonicWalletUnit_LoadAddressByPath(t *testing.T) {
	type testCase struct {
		Mnemonic    string
		AddressPath *pbCommon.DerivationAddressIdentity

		ExpectedAddress string
	}

	// WARN: DO NOT USE THESE MNEMONICS IN MAINNET OR TESTNET. Usage only in unit-tests
	// WARN: DO NOT USE THESE MNEMONICS IN MAINNET OR TESTNET. Usage only in unit-tests
	// WARN: DO NOT USE THESE MNEMONICS IN MAINNET OR TESTNET. Usage only in unit-tests
	testCases := []*testCase{
		{
			Mnemonic: "umbrella uphold security hill monkey skin either immense kid afraid sense desk extend twenty doctor odor buzz reject derive frame hub much once suffer",
			AddressPath: &pbCommon.DerivationAddressIdentity{
				AccountIndex:  5,
				InternalIndex: 12,
				AddressIndex:  3,
			},
			ExpectedAddress: "0x10610cABbc97DD5ccd46B2571eE9a7968822A63C",
		},
		{
			Mnemonic: "slogan follow oil world head protect patrol wagon toddler fly kangaroo kite dash essay shoulder worth one grace shift good disease biology magic pottery",
			AddressPath: &pbCommon.DerivationAddressIdentity{
				AccountIndex:  1000,
				InternalIndex: 10000,
				AddressIndex:  100000,
			},
			ExpectedAddress: "0xEfFB3DBe9e1a893ee2e90B93e067E266FEf3196a",
		},
		{
			Mnemonic: "image video differ dumb later child gather smart supply mountain salon ring boy mystery hope secret present bar then joke latin guitar view devote",
			AddressPath: &pbCommon.DerivationAddressIdentity{
				AccountIndex:  1,
				InternalIndex: 102,
				AddressIndex:  1003,
			},
			ExpectedAddress: "0x3ca3d79305Ee17c2707cF5114b5Be30e5f88eB3F",
		},
	}

	for _, tCase := range testCases {
		poolUnitIntrf, err := NewPoolUnit(uuid.NewString(), tCase.Mnemonic)
		if err != nil {
			t.Fatalf("%s: %e", "unable to create mnemonic wallet pool unit:", err)
		}

		poolUnit, ok := poolUnitIntrf.(*mnemonicWalletUnit)
		if !ok {
			t.Fatalf("%s", "unable to cast interface to pool unit worker")
		}

		accountIdentity := &anypb.Any{}
		_ = accountIdentity.MarshalFrom(tCase.AddressPath)

		addr, err := poolUnit.LoadAccount(context.Background(), accountIdentity)
		if err != nil {
			t.Fatalf("%s: %e", "unable to get address from pool unit:", err)
		}

		if addr == nil {
			t.Fatalf("%s: %e", "missing address in pool unit result:", err)
		}

		if len(poolUnit.addressPool) == 0 {
			t.Fatalf("%s", "address in pool not loaded")
		}

		key := fmt.Sprintf(addrPatKeyTemplate, tCase.AddressPath.AccountIndex,
			tCase.AddressPath.InternalIndex, tCase.AddressPath.AddressIndex)
		addrData, ok := poolUnit.addressPool[key]
		if !ok || addrData == nil {
			t.Fatalf("%s", "missing data by key in address pool")
		}

		if addrData.privateKey == nil {
			t.Fatalf("%s", "missing private key in address pool unit")
		}

		if tCase.ExpectedAddress != addrData.address {
			t.Fatalf("%s", "address not equal with expected")
		}

		if tCase.ExpectedAddress != *addr {
			t.Fatalf("%s", "address not equal with expected")
		}
	}
}

func TestMnemonicWalletUnit_SignData(t *testing.T) {
	type testCase struct {
		Mnemonic    string
		AddressPath *pbCommon.DerivationAddressIdentity

		CompressedPublicKey string
		PublicKey           string

		DataForSign *pbEthereum.LegacyTxData

		ExpectedAddress    string
		ExpectedSignedData []byte
	}

	// WARN: DO NOT USE THESE MNEMONICS IN MAINNET OR TESTNET. Usage only in unit-tests
	// WARN: DO NOT USE THESE MNEMONICS IN MAINNET OR TESTNET. Usage only in unit-tests
	// WARN: DO NOT USE THESE MNEMONICS IN MAINNET OR TESTNET. Usage only in unit-tests
	testCases := []*testCase{
		{
			Mnemonic: "unknown valid carbon hat echo funny artist letter desk absorb unit fatigue foil skirt stay case path rescue hawk remember aware arch regular cry",
			AddressPath: &pbCommon.DerivationAddressIdentity{
				AccountIndex:  7,
				InternalIndex: 8,
				AddressIndex:  9,
			},
			DataForSign: &pbEthereum.LegacyTxData{
				Nonce:          0,
				GasPrice:       big.NewInt(10000000000).Bytes(),
				GasLimit:       14000,
				ToAddress:      common.HexToAddress("0xBE0eB53F46cd790Cd13851d5EFf43D12404d33E8").Bytes(),
				Value:          big.NewInt(1500000).Bytes(),
				Data:           nil,
				SignParameters: nil,
			},

			CompressedPublicKey: "0x030dbc0361bb42cedfce71de5bd969a573d952e23e44a905bd89d44e3b3b2bafb0",
			PublicKey:           "0x040dbc0361bb42cedfce71de5bd969a573d952e23e44a905bd89d44e3b3b2bafb08142c1ace72356a45089dab3429746a68182c5398851b5baa835375560d2755b",
			ExpectedAddress:     "0xf8A0F16782625B16260D0A4b0Ed107412bd95d56",
		},
		{
			Mnemonic: "laundry file mystery rate absorb wrist despair cook near afraid account mirror name chair lake regular vicious oblige release vicious identify glimpse flight help",
			AddressPath: &pbCommon.DerivationAddressIdentity{
				AccountIndex:  909,
				InternalIndex: 8008,
				AddressIndex:  70007,
			},
			DataForSign: &pbEthereum.LegacyTxData{
				Nonce:          12,
				GasPrice:       big.NewInt(10000200000).Bytes(),
				GasLimit:       15000,
				ToAddress:      common.HexToAddress("0x40B38765696e3d5d8d9d834D8AaD4bB6e418E489").Bytes(),
				Value:          big.NewInt(1900000).Bytes(),
				Data:           []byte{0x1, 0x2, 0x3},
				SignParameters: nil,
			},

			CompressedPublicKey: "0x0370cd063d2aa795777bd6111e8f0ab6bc062a91721dbf9018cba3ab0e7eb3d798",
			PublicKey:           "0x0470cd063d2aa795777bd6111e8f0ab6bc062a91721dbf9018cba3ab0e7eb3d79868d853f137d5d4096b715a7d480590bde3ee6cd42beab9b638eeeec8ec1da08b",
			ExpectedAddress:     "0x9C4Bb7A12cAd7145682854BBCBF9CaDfFD4EEc01",
		},
		{
			Mnemonic: "busy spawn solar december element round wild buddy furnace help clog tired object camera resist maze fuel need stock rule spot diagram aisle expect",
			AddressPath: &pbCommon.DerivationAddressIdentity{
				AccountIndex:  9,
				InternalIndex: 8,
				AddressIndex:  7,
			},
			CompressedPublicKey: "0x027465728e9c06c6c2da32e96159ae8e15ec1381baba99c7be2607dc6604594830",
			PublicKey:           "0x047465728e9c06c6c2da32e96159ae8e15ec1381baba99c7be2607dc6604594830e58aa45697f4ff1dd1ea64780da8c2217a2b5ca0a6f35e47f58877d3ddc44938",
			DataForSign: &pbEthereum.LegacyTxData{
				Nonce:          12,
				GasPrice:       big.NewInt(10000200000).Bytes(),
				GasLimit:       15000,
				ToAddress:      common.HexToAddress("0x61EDCDf5bb737ADffE5043706e7C5bb1f1a56eEA").Bytes(),
				Value:          big.NewInt(1900000).Bytes(),
				Data:           []byte{0x9, 0x10, 0x11, 0x12},
				SignParameters: nil,
			},

			ExpectedAddress: "0xdD4aE0268C7F6144bDb08C1dEC349bCe5f239E30",
		},
	}

	for _, tCase := range testCases {
		poolUnitIntrf, loopErr := NewPoolUnit(uuid.NewString(), tCase.Mnemonic)
		if loopErr != nil {
			t.Fatalf("%s: %e", "unable to create mnemonic wallet pool unit:", loopErr)
		}

		poolUnit, ok := poolUnitIntrf.(*mnemonicWalletUnit)
		if !ok {
			t.Fatalf("%s", "unable to cast interface to pool unit worker")
		}

		accountIdentity := &anypb.Any{}
		_ = accountIdentity.MarshalFrom(tCase.AddressPath)

		signDataAny := &anypb.Any{}
		_ = signDataAny.MarshalFrom(tCase.DataForSign)

		signDataRaw, loopErr := proto.Marshal(signDataAny)
		if loopErr != nil {
			t.Fatalf("%s: %e", "unable to marshal proto signature proto message:", loopErr)
		}

		addr, signedData, loopErr := poolUnit.SignData(context.Background(), accountIdentity, signDataRaw)
		if loopErr != nil {
			t.Fatalf("%s: %e", "unable to sign data:", loopErr)
		}

		if addr == nil {
			t.Fatalf("%s", "missing address in result of sign method")
		}

		if signedData == nil {
			t.Fatalf("%s", "missing signed data in result of sign method")
		}

		if len(poolUnit.addressPool) == 0 {
			t.Fatalf("%s", "address in pool not loaded")
		}

		key := fmt.Sprintf(addrPatKeyTemplate, tCase.AddressPath.AccountIndex,
			tCase.AddressPath.InternalIndex, tCase.AddressPath.AddressIndex)
		addrData, ok := poolUnit.addressPool[key]
		if !ok || addrData == nil {
			t.Fatalf("%s", "missing data by key in address pool")
		}

		if addrData.privateKey == nil {
			t.Fatalf("%s", "missing private key in address pool unit")
		}

		if tCase.ExpectedAddress != addrData.address {
			t.Fatalf("%s", "address not equal with expected")
		}

		if tCase.ExpectedAddress != *addr {
			t.Fatalf("%s", "address not equal with expected")
		}

		signed := bytes.Clone(signedData)

		//h256h := sha256.New()
		//h256h.Write(tCase.DataForSign)
		//hash := h256h.Sum(nil)
		//
		//h256h.Reset()
		//h256h = nil

		tx := &types.Transaction{}
		loopErr = tx.UnmarshalBinary(signed)
		if loopErr != nil {
			t.Fatalf("%s: %e", "unable to unmarshal tx data", loopErr)
		}
		//
		//arr := func(sig []byte) (r, s, v *big.Int) {
		//	if len(sig) != crypto.SignatureLength {
		//		panic(fmt.Sprintf("wrong size for signature: got %d, want %d", len(sig), crypto.SignatureLength))
		//	}
		//	r = new(big.Int).SetBytes(sig[:32])
		//	s = new(big.Int).SetBytes(sig[32:64])
		//	v = new(big.Int).SetBytes([]byte{sig[64] + 27})
		//	return r, s, v
		//}

		//hash := tx.Hash()
		log.Print(tx.Hash().String())

		recPubKey, recAddr, loopErr := extractECSDAPublicKey(tx)
		if loopErr != nil {
			t.Errorf("%s: %e", "unable to extract publick key and address", loopErr)
		}

		//sigPublicKeyECDSA, loopErr := crypto.SigToPub(hash.Bytes()[:], sig)
		//if loopErr != nil {
		//	t.Errorf("%s: %e", "unable to get public key from signed message", loopErr)
		//}

		sigPublicKeyECDSABytes := crypto.FromECDSAPub(recPubKey)
		sigPublicKeyECDSAString := hexutil.Encode(sigPublicKeyECDSABytes)

		if tCase.PublicKey != sigPublicKeyECDSAString {
			t.Errorf("%s", "ethereumWallet pubKey not equal with expected")
		}

		//sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), sig)
		//if err != nil {
		//	t.Error(err)
		//}
		//sigPubKeyString := hexutil.Encode(sigPublicKey)
		//
		//if tCase.PublicKey != sigPubKeyString {
		//	t.Errorf("%s", "ethereumWallet pubKey not equal with expected")
		//}

		compressed := crypto.CompressPubkey(recPubKey)
		compressedSigPublicKeyECDSAString := hexutil.Encode(compressed)
		if tCase.CompressedPublicKey != compressedSigPublicKeyECDSAString {
			t.Errorf("%s", "ethereumWallet addr from pubKey not equal with expected")
		}

		ethAddr := crypto.PubkeyToAddress(*recPubKey)
		ethAddrStr := ethAddr.String()

		recAddrString := recAddr.String()
		if tCase.ExpectedAddress != recAddrString {
			t.Errorf("%s", "ethereumWallet addr from pubKey not equal with expected")
		}

		if tCase.ExpectedAddress != ethAddrStr {
			t.Errorf("%s", "ethereumWallet addr from pubKey not equal with expected")
		}

		if tCase.ExpectedAddress != *addr {
			t.Errorf("%s", "ethereumWallet addr from pubKey not equal with expected")
		}
	}
}

func TestMnemonicWalletUnit_UnloadWallet(t *testing.T) {
	type testCase struct {
		Mnemonic    string
		AddressPath *pbCommon.DerivationAddressIdentity

		ExpectedAddress string
	}

	// WARN: DO NOT USE THESE MNEMONICS IN MAINNET OR TESTNET. Usage only in unit-tests
	// WARN: DO NOT USE THESE MNEMONICS IN MAINNET OR TESTNET. Usage only in unit-tests
	// WARN: DO NOT USE THESE MNEMONICS IN MAINNET OR TESTNET. Usage only in unit-tests
	testCases := []*testCase{
		{
			Mnemonic: "input erase buzz crew miss auction habit cargo wrestle perfect like midnight buddy chase grit only treat stuff rival worth alien tennis parent artist",
			AddressPath: &pbCommon.DerivationAddressIdentity{
				AccountIndex:  5,
				InternalIndex: 8,
				AddressIndex:  11,
			},
			ExpectedAddress: "0xb0160385Da6536fb6790466Fc9Eef8F9C1384A25",
		},
		{
			Mnemonic: "empower plate axis divorce neither noodle above flight very indoor zone mango sand exhaust nominee solid combine picnic gospel myth stem raw garage veteran",
			AddressPath: &pbCommon.DerivationAddressIdentity{
				AccountIndex:  2,
				InternalIndex: 4,
				AddressIndex:  8,
			},
			ExpectedAddress: "0x05843f146f4F0186b27983E87342Eb17A27F111F",
		},
		{
			Mnemonic: "sea vault tattoo laugh ugly where saddle six usage install one cube affair sick used actress zebra fuel sunny tackle can siege develop drop",
			AddressPath: &pbCommon.DerivationAddressIdentity{
				AccountIndex:  8,
				InternalIndex: 64,
				AddressIndex:  4096,
			},
			ExpectedAddress: "0xf731F8246F965A3CaC805D56ca809586B9786eF3",
		},
	}

	for _, tCase := range testCases {
		poolUnitIntrf, loopErr := NewPoolUnit(uuid.NewString(), tCase.Mnemonic)
		if loopErr != nil {
			t.Fatalf("%s: %e", "unable to create mnemonic wallet pool unit:", loopErr)
		}

		poolUnit, ok := poolUnitIntrf.(*mnemonicWalletUnit)
		if !ok {
			t.Fatalf("%s", "unable to cast interface to pool unit worker")
		}

		accountIdentity := &anypb.Any{}
		_ = accountIdentity.MarshalFrom(tCase.AddressPath)

		addr, loopErr := poolUnit.LoadAccount(context.Background(), accountIdentity)
		if loopErr != nil {
			t.Fatalf("%s: %e", "unable to sign data:", loopErr)
		}

		if addr == nil {
			t.Fatalf("%s", "missing address in result of sign method")
		}

		if len(poolUnit.addressPool) == 0 {
			t.Fatalf("%s", "address in pool not loaded")
		}

		key := fmt.Sprintf(addrPatKeyTemplate, tCase.AddressPath.AccountIndex,
			tCase.AddressPath.InternalIndex, tCase.AddressPath.AddressIndex)
		addrData, ok := poolUnit.addressPool[key]
		if !ok || addrData == nil {
			t.Fatalf("%s", "missing data by key in address pool")
		}

		loopErr = poolUnit.UnloadWallet()
		if loopErr != nil {
			t.Fatalf("%s: %e", "unable to unload wallet", loopErr)
		}

		if len(poolUnit.addressPool) != 0 {
			t.Fatalf("%s", "address pool is not empty")
		}

		if poolUnit.hdWalletSvc != nil {
			t.Fatalf("%s", "hdwallet service is not nil")
		}

		if poolUnit.mnemonicHash != "0" {
			t.Fatalf("%s", "mnemonicHash is not equal zero")
		}
	}
}
