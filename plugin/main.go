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
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/core/types"
)

const (
	evmDefaultPluginName   = "ethereum-hdwallet-plugin"
	ethereumMainNetChainID = 1
)

// DO NOT EDIT THESE VARIABLES DIRECTLY. These are build-time constants
// DO NOT USE THESE VARIABLES IN APPLICATION CODE. USE commonConfig.NewLdFlagsManager SERVICE-COMPONENT INSTEAD OF IT
var (
	// ReleaseTag - release tag in TAG.SHORT_COMMIT_ID.BUILD_NUMBER.
	// DO NOT EDIT THIS VARIABLE DIRECTLY. These are build-time constants
	// DO NOT USE THESE VARIABLES IN APPLICATION CODE
	ReleaseTag = "v0.0.0-00000000-100500"

	// CommitID - latest commit id.
	// DO NOT EDIT THIS VARIABLE DIRECTLY. These are build-time constants
	// DO NOT USE THESE VARIABLES IN APPLICATION CODE
	CommitID = "0000000000000000000000000000000000000000"

	// ShortCommitID - first 12 characters from CommitID.
	// DO NOT EDIT THIS VARIABLE DIRECTLY. These are build-time constants
	// DO NOT USE THESE VARIABLES IN APPLICATION CODE
	ShortCommitID = "0000000"

	// BuildNumber - ci/cd build number for BuildNumber
	// DO NOT EDIT THIS VARIABLE DIRECTLY. These are build-time constants
	// DO NOT USE THESE VARIABLES IN APPLICATION CODE
	BuildNumber string = "100500"

	// BuildDateTS - ci/cd build date in time stamp
	// DO NOT EDIT THIS VARIABLE DIRECTLY. These are build-time constants
	// DO NOT USE THESE VARIABLES IN APPLICATION CODE
	BuildDateTS string = "1713280105"

	// NetworkChainID - blockchain network ID, Ethereum blockchain mainnet id = 1
	// https://chainlist.org/
	// https://docs.expand.network/important-ids/chain-id
	// https://chainid.network/chains.json - json format
	// DO NOT EDIT THIS VARIABLE DIRECTLY. These are build-time constants
	// DO NOT USE THESE VARIABLES IN APPLICATION CODE
	NetworkChainID = "1"

	// CoinType - registered coin type from BIP-0044 standard.
	// Default value for Ethereum = 60, Ethereum Classic = 61, Binance Smart Chain = 9006, etc...
	// See BIP-0044 - https://github.com/satoshilabs/slips/blob/master/slip-0044.md
	// But typically for all EVM blockchains use default Ethereum value = 60
	// DO NOT EDIT THIS VARIABLE DIRECTLY. These are build-time constants
	// DO NOT USE THESE VARIABLES IN APPLICATION CODE
	CoinType = "60"

	// NetworkName - name of the network for which the plugin was built.
	// Default value Ethereum MainNet network name
	// DO NOT EDIT THIS VARIABLE DIRECTLY. These are build-time constants
	// DO NOT USE THESE VARIABLES IN APPLICATION CODE
	NetworkName = evmDefaultPluginName
)

var (
	pluginChainID  = ethereumMainNetChainID
	pluginCoinType = ethereumCoinNumber
	pluginName     = evmDefaultPluginName
	pluginSigner   types.Signer

	prepareChainIDOnce       = sync.Once{}
	setChainIDOnce           = sync.Once{}
	prepareCoinTypeOnce      = sync.Once{}
	setCoinTypeOnce          = sync.Once{}
	setSignerOnce            = sync.Once{}
	setPluginNetworkNameOnce = sync.Once{}

	ErrPluginValueAlreadySet = errors.New("plugin value already set.You can do it only once")
)

func init() {
	preparePluginNetworkName()
	prepareChainID()
	prepareCoinType()
	prepareSigner()
}

func GetPluginName() string {
	return pluginName
}

func GetPluginReleaseTag() string {
	return ReleaseTag
}

func GetPluginCommitID() string {
	return CommitID
}

func GetPluginShortCommitID() string {
	return ShortCommitID
}

func GetPluginBuildNumber() string {
	return BuildNumber
}

func GetPluginBuildDateTS() string {
	return BuildDateTS
}

func GetChainID() int {
	return pluginChainID
}

func GetSupportedChainIDsInfo() string {
	return "Plugin support all EVM-like blockchains.\n" +
		"All you need it set right values of NetworkChainID variable.\n" +
		"Please reade README.md file for getting information about available flows"
}

func SetChainID(chainID int) error {
	return setChainID(chainID)
}

func GetSupportedCoinTypesInfo() string {
	return "Plugin support all EVM-like blockchains.\n" +
		"All you need it set right values of CoinType variables.\n" +
		"Please reade README.md file for getting information about available flows"
}

func GetHdWalletCoinType() int {
	return pluginCoinType
}

func SetHdWalletCoinType(coinType int) error {
	return setHdWalletCoinType(coinType)
}

func preparePluginNetworkName() string {
	setPluginNetworkNameOnce.Do(func() {
		if NetworkName == "" {
			pluginName = evmDefaultPluginName

			return
		}

		pluginName = NetworkName

		return
	})

	return pluginName
}

func setChainID(chainID int) error {
	var err = fmt.Errorf("%w: %s", ErrPluginValueAlreadySet, "pluginChainID")
	setChainIDOnce.Do(func() {
		pluginChainID = chainID
		err = nil

		return
	})

	return err
}

func setHdWalletCoinType(coinType int) error {
	var err = fmt.Errorf("%w: %s", ErrPluginValueAlreadySet, "pluginCoinType")
	setCoinTypeOnce.Do(func() {
		pluginCoinType = coinType
		err = nil

		return
	})

	return err
}

func prepareChainID() int {
	prepareChainIDOnce.Do(func() {
		if NetworkChainID == "" {
			pluginCoinType = ethereumMainNetChainID

			return
		}

		chainIDInt, err := strconv.Atoi(NetworkChainID)
		if err != nil {
			panic(fmt.Errorf("wrong chain id type format: %w", err))
		}

		pluginChainID = chainIDInt

		return
	})

	return pluginChainID
}

func prepareCoinType() int {
	prepareCoinTypeOnce.Do(func() {
		if CoinType == "" {
			pluginCoinType = ethereumCoinNumber

			return
		}

		coinTypeInt, err := strconv.Atoi(CoinType)
		if err != nil {
			panic(fmt.Errorf("wrong coin type format: %w", err))
		}

		pluginCoinType = coinTypeInt

		return
	})

	return pluginCoinType
}

func prepareSigner() types.Signer {
	setSignerOnce.Do(func() {
		pluginSigner = types.LatestSignerForChainID(big.NewInt(int64(pluginChainID)))
	})

	return pluginSigner
}
