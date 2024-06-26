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
	"fmt"
	pbCommon "github.com/crypto-bundle/bc-wallet-common-hdwallet-controller/pkg/grpc/common"
	"github.com/google/uuid"
	"github.com/tyler-smith/go-bip39"
	"google.golang.org/protobuf/types/known/anypb"
	"log"
	"os"
	"plugin"
	"strconv"
	"time"
)

type walletPoolUnitService interface {
	UnloadWallet() error

	GetWalletUUID() string
	LoadAccount(ctx context.Context,
		accountParameters *anypb.Any,
	) (*string, error)
	GetAccountAddress(ctx context.Context,
		accountParameters *anypb.Any,
	) (*string, error)
	GetMultipleAccounts(ctx context.Context,
		multipleAccountsParameters *anypb.Any,
	) (uint, []*pbCommon.AccountIdentity, error)
	SignData(ctx context.Context,
		accountParameters *anypb.Any,
		dataForSign []byte,
	) (*string, []byte, error)
}

const (
	ethereumMainNetCoinType = 60
	ethereumMainNetChainID  = 1

	getPluginNameSymbol          = "GetPluginName"
	getPluginReleaseTagSymbol    = "GetPluginReleaseTag"
	getPluginCommitIDSymbol      = "GetPluginCommitID"
	getPluginShortCommitIDSymbol = "GetPluginShortCommitID"
	getPluginBuildNumberSymbol   = "GetPluginBuildNumber"
	getPluginBuildDateTSSymbol   = "GetPluginBuildDateTS"

	pluginGetChainIDSymbol               = "GetChainID"
	pluginGetSupportedChainIDsInfoSymbol = "GetSupportedChainIDsInfo"
	pluginSetChainIDSymbol               = "SetChainID"

	pluginGetCoinTypeSymbol               = "GetHdWalletCoinType"
	pluginGetSupportedCoinTypesInfoSymbol = "GetSupportedCoinTypesInfo"
	pluginSetCoinTypeSymbol               = "SetHdWalletCoinType"

	pluginGenerateMnemonicSymbol = "GenerateMnemonic"
	pluginValidateMnemonicSymbol = "ValidateMnemonic"
	pluginNewPoolUnitSymbol      = "NewPoolUnit"
)

var stringFuncSymbolLookUp = func(plugin *plugin.Plugin, symbolName string) (func() string, error) {
	pluginFuncSymbol, err := plugin.Lookup(symbolName)
	if err != nil {
		return nil, err
	}

	pluginFunc, ok := pluginFuncSymbol.(func() string)
	if !ok {
		return nil, fmt.Errorf("%s: %s", "unable to cast plugin symbol", symbolName)
	}

	return pluginFunc, nil
}

func main() {

	p, err := plugin.Open("./build/ethereum.so")
	if err != nil {
		log.Fatalf("%s: %e", "unable to load pluggable extension", err)
	}

	runGetPluginNameTest(p)
	runReleaseTagTest(p)

	runGetChainIdTest(p)
	runGetSupportedChainIDsTest(p)
	runSetChainIdTest(p)

	runGetCoinTypeTest(p)
	runGetSupportedCoinTypesInfoTest(p)
	runSetCoinTypeTest(p)

	runGetCommitIDTest(p)
	runGetShortCommitIDTest(p)
	runGetPluginBuildNumberTest(p)
	runGetPluginBuildBuildDateTest(p)
	runGenerateMnemonicTest(p)
	runValidateMnemonicTest(p)
	runNewWalletPoolTest(p)

	log.Println("PASS")

	os.Exit(0)
}

func runGetPluginNameTest(p *plugin.Plugin) {
	log.Printf("=== RUN: %s\n", getPluginNameSymbol)

	getPluginNameFunc, err := stringFuncSymbolLookUp(p, getPluginNameSymbol)
	if err != nil {
		log.Fatal(err)
	}

	if getPluginNameFunc == nil {
		log.Fatal("missing Get release tag function")
	}

	pluginName := getPluginNameFunc()
	if len(pluginName) == 0 {
		log.Fatalf("%s: %d", "zero length of release tag value", len(pluginName))
	}

	log.Printf("--- PASS: %s\n", getPluginNameSymbol)
}

func runReleaseTagTest(p *plugin.Plugin) {
	log.Printf("=== RUN: %s\n", getPluginReleaseTagSymbol)

	getPluginReleaseTagFunc, err := stringFuncSymbolLookUp(p, getPluginReleaseTagSymbol)
	if err != nil {
		log.Fatal(err)
	}

	if getPluginReleaseTagFunc == nil {
		log.Fatal("missing Get release tag function")
	}

	releaseTag := getPluginReleaseTagFunc()
	if len(releaseTag) == 0 {
		log.Fatalf("%s: %d", "zero length of release tag value", len(releaseTag))
	}

	log.Printf("--- PASS: %s\n", getPluginReleaseTagSymbol)
}

func runGetCoinTypeTest(p *plugin.Plugin) {
	log.Printf("=== RUN: %s\n", pluginGetCoinTypeSymbol)

	pluginGetCoinTypeFuncSymbol, err := p.Lookup(pluginGetCoinTypeSymbol)
	if err != nil {
		log.Fatal(err)
	}

	getCoinTypeFunc, ok := pluginGetCoinTypeFuncSymbol.(func() int)
	if !ok {
		log.Fatalf("%s: %s", "unable to cast plugin symbol", pluginGetCoinTypeSymbol)
	}

	currentCoinType := getCoinTypeFunc()
	if currentCoinType <= 0 {
		log.Fatalf("%s - current: %d", "wrong chainID value", currentCoinType)
	}

	log.Printf("--- PASS: %s\n", pluginGetCoinTypeSymbol)
}

func runGetSupportedCoinTypesInfoTest(p *plugin.Plugin) {
	log.Printf("=== RUN: %s\n", pluginGetSupportedCoinTypesInfoSymbol)

	supportedCoinIDsPluginFuncSymbol, err := p.Lookup(pluginGetSupportedCoinTypesInfoSymbol)
	if err != nil {
		log.Fatal(err)
	}

	getSupportedChainIDsFunc, ok := supportedCoinIDsPluginFuncSymbol.(func() string)
	if !ok {
		log.Fatalf("%s: %s", "unable to cast plugin symbol", pluginGetSupportedCoinTypesInfoSymbol)
	}

	chainIDList := getSupportedChainIDsFunc()
	if len(chainIDList) == 0 {
		log.Fatalf("%s: %d", "empty supported coinID info", len(chainIDList))
	}

	log.Printf("--- PASS: %s\n", pluginGetSupportedCoinTypesInfoSymbol)
}

func runSetCoinTypeTest(p *plugin.Plugin) {
	log.Printf("=== RUN: %s\n", pluginSetCoinTypeSymbol)

	pluginSetCoinTypeFuncSymbol, err := p.Lookup(pluginSetCoinTypeSymbol)
	if err != nil {
		log.Fatal(err)
	}

	setCoinTypeFunc, ok := pluginSetCoinTypeFuncSymbol.(func(int) error)
	if !ok {
		log.Fatalf("%s: %s", "unable to cast plugin symbol", pluginSetCoinTypeSymbol)
	}

	err = setCoinTypeFunc(ethereumMainNetCoinType)
	if err != nil {
		log.Fatalf("--- FAIL: %s, unable to set chainID - expected: %d. Error: %s", pluginSetCoinTypeSymbol,
			ethereumMainNetCoinType, err)
	}

	err = setCoinTypeFunc(ethereumMainNetCoinType + 1)
	if err != nil {
		log.Printf("--- INFO: %s, unable to set coinType second time. Its OK. Error: %s",
			pluginSetCoinTypeSymbol, err)
	}

	log.Printf("--- PASS: %s\n", pluginSetCoinTypeSymbol)
}

func runSetChainIdTest(p *plugin.Plugin) {
	log.Printf("=== RUN: %s\n", pluginSetChainIDSymbol)

	pluginSetChainIDFuncSymbol, err := p.Lookup(pluginSetChainIDSymbol)
	if err != nil {
		log.Fatal(err)
	}

	setChainIDFunc, ok := pluginSetChainIDFuncSymbol.(func(int) error)
	if !ok {
		log.Fatalf("%s: %s", "unable to cast plugin symbol", pluginSetChainIDSymbol)
	}

	err = setChainIDFunc(ethereumMainNetChainID)
	if err != nil {
		log.Fatalf("%s - expected: %d %d", "unable to set chainID", 199)
	}

	err = setChainIDFunc(ethereumMainNetChainID + 1)
	if err != nil {
		log.Printf("--- INFO: %s, unable to set chainID second time. Its OK. Error: %s",
			pluginSetCoinTypeSymbol, err)
	}

	log.Printf("--- PASS: %s\n", pluginSetChainIDSymbol)
}

func runGetChainIdTest(p *plugin.Plugin) {
	log.Printf("=== RUN: %s\n", pluginGetChainIDSymbol)

	pluginGetCoinIDFuncSymbol, err := p.Lookup(pluginGetChainIDSymbol)
	if err != nil {
		log.Fatal(err)
	}

	getCoinIDFunc, ok := pluginGetCoinIDFuncSymbol.(func() int)
	if !ok {
		log.Fatalf("%s: %s", "unable to cast plugin symbol", pluginGetChainIDSymbol)
	}

	currentChainID := getCoinIDFunc()
	if currentChainID != ethereumMainNetChainID {
		log.Fatalf("%s - expected: %d, current: %d", "wrong chainID value",
			ethereumMainNetChainID, currentChainID)
	}

	log.Printf("--- PASS: %s\n", pluginGetChainIDSymbol)
}

func runGetSupportedChainIDsTest(p *plugin.Plugin) {
	log.Printf("=== RUN: %s\n", pluginGetSupportedChainIDsInfoSymbol)

	supportedCoinIDsPluginFuncSymbol, err := p.Lookup(pluginGetSupportedChainIDsInfoSymbol)
	if err != nil {
		log.Fatal(err)
	}

	getSupportedChainIDsFunc, ok := supportedCoinIDsPluginFuncSymbol.(func() string)
	if !ok {
		log.Fatalf("%s: %s", "unable to cast plugin symbol", pluginGetSupportedChainIDsInfoSymbol)
	}

	chainIDList := getSupportedChainIDsFunc()
	if len(chainIDList) == 0 {
		log.Fatalf("%s: %d", "empty supported coinID list", len(chainIDList))
	}

	log.Printf("--- PASS: %s\n", pluginGetSupportedChainIDsInfoSymbol)
}

func runGetCommitIDTest(p *plugin.Plugin) {
	log.Printf("=== RUN: %s\n", getPluginCommitIDSymbol)

	getPluginCommitIDFunc, err := stringFuncSymbolLookUp(p, getPluginCommitIDSymbol)
	if err != nil {
		log.Fatal(err)
	}

	if getPluginCommitIDFunc == nil {
		log.Fatalf("missing Get commit id function")
	}

	commitID := getPluginCommitIDFunc()
	if len(commitID) != 40 {
		log.Fatalf("%s: %d, %s", "wrong length of commit ID value", len(commitID), commitID)
	}

	log.Printf("--- PASS: %s\n", getPluginCommitIDSymbol)
}

func runGetShortCommitIDTest(p *plugin.Plugin) {
	log.Printf("=== RUN: %s\n", getPluginShortCommitIDSymbol)

	getPluginShortCommitIDFunc, err := stringFuncSymbolLookUp(p, getPluginShortCommitIDSymbol)
	if err != nil {
		log.Fatal(err)
	}

	if getPluginShortCommitIDFunc == nil {
		log.Fatal("missing Get short commit id function")
	}

	shortCommitID := getPluginShortCommitIDFunc()
	if len(shortCommitID) != 7 {
		log.Fatalf("%s: %d, %s", "wrong length of short commit ID value", len(shortCommitID),
			shortCommitID)
	}

	log.Printf("--- PASS: %s\n", getPluginShortCommitIDSymbol)
}

func runGetPluginBuildNumberTest(p *plugin.Plugin) {
	log.Printf("=== RUN: %s\n", getPluginBuildNumberSymbol)

	getPluginBuildNumberFunc, err := stringFuncSymbolLookUp(p, getPluginBuildNumberSymbol)
	if err != nil {
		log.Fatal(err)
	}

	if getPluginBuildNumberFunc == nil {
		log.Fatal("missing Get plugin build number function")
	}

	buildNumber := getPluginBuildNumberFunc()
	if _, err = strconv.ParseUint(buildNumber, 10, 0); err != nil {
		log.Fatalf("%s: %e", "wrong build number format", err)
	}

	log.Printf("--- PASS: %s\n", getPluginBuildNumberSymbol)
}

func runGetPluginBuildBuildDateTest(p *plugin.Plugin) {
	log.Printf("=== RUN: %s\n", getPluginBuildDateTSSymbol)

	getPluginBuildDateTSFunc, err := stringFuncSymbolLookUp(p, getPluginBuildDateTSSymbol)
	if err != nil {
		log.Fatal(err)
	}

	if getPluginBuildDateTSFunc == nil {
		log.Fatal("missing Get plugin build date function")
	}

	buildDate := getPluginBuildDateTSFunc()
	buildDataTS, err := strconv.ParseUint(buildDate, 10, 0)
	if err != nil {
		log.Fatalf("%s: %e", "wrong build date Time stamp format", err)
	}
	tm := time.Unix(int64(buildDataTS), 0)
	tmString := strconv.FormatUint(uint64(tm.Unix()), 10)
	if buildDate != tmString {
		log.Fatalf("%s, current value: %s", "something wrong with build date time stamp string",
			buildDate)
	}

	log.Printf("--- PASS: %s\n", getPluginBuildDateTSSymbol)
}

func runGenerateMnemonicTest(p *plugin.Plugin) {
	log.Printf("=== RUN: %s\n", pluginGenerateMnemonicSymbol)

	generateMnemonicFuncSymbol, err := p.Lookup(pluginGenerateMnemonicSymbol)
	if err != nil {
		log.Fatal(err)
	}

	if generateMnemonicFuncSymbol == nil {
		log.Fatal("missing Generate mnemonic function symbol")
	}

	generateMnemoFunc, ok := generateMnemonicFuncSymbol.(func() (string, error))
	if !ok {
		log.Fatal("unable to cast generate mnemonic function")
	}

	generatedMnemonic, err := generateMnemoFunc()
	if err != nil {
		log.Fatal(err)
	}

	if !bip39.IsMnemonicValid(generatedMnemonic) {
		log.Fatal("generated mnemonic phares is not valid")
	}

	log.Printf("--- PASS: %s\n", pluginGenerateMnemonicSymbol)
}

func runValidateMnemonicTest(p *plugin.Plugin) {
	log.Printf("=== RUN: %s\n", pluginValidateMnemonicSymbol)

	validateMnemonicFuncSymbol, err := p.Lookup(pluginValidateMnemonicSymbol)
	if err != nil {
		log.Fatal(err)
	}

	if validateMnemonicFuncSymbol == nil {
		log.Fatal("missing Validate mnemonic function symbol")
	}

	validateMnemoFunc, isCasted := validateMnemonicFuncSymbol.(func(mnemonic string) bool)
	if !isCasted {
		log.Fatal("unable to cast validate function")
	}

	if validateMnemoFunc == nil {
		log.Fatal("missing Validate mnemonic function")
	}

	// WARN: DO NOT USE THESE MNEMONICS IN MAINNET OR TESTNET. Usage only in unit-tests
	mnemoPhrase := "beach large spray gentle buyer hover flock dream hybrid match whip ten mountain pitch enemy lobster afford barrel patrol desk trigger output excuse truck"
	if !validateMnemoFunc(mnemoPhrase) {
		log.Fatal("failed mnemonic validation validation")
	}

	log.Printf("--- PASS: %s\n", pluginValidateMnemonicSymbol)
}

func runNewWalletPoolTest(p *plugin.Plugin) {
	log.Printf("=== RUN: %s\n", pluginNewPoolUnitSymbol)

	unitMakerFuncSymbol, err := p.Lookup(pluginNewPoolUnitSymbol)
	if err != nil {
		log.Fatal(err)
	}

	unitMakerFunc, isCasted := unitMakerFuncSymbol.(func(walletUUID string,
		mnemonicDecryptedData string,
	) (interface{}, error))
	if !isCasted {
		log.Fatal("unable to cast pool unit Maker function")
	}

	// WARN: DO NOT USE THESE MNEMONICS IN MAINNET OR TESTNET. Usage only in unit-tests
	mnemoPhrase := "beach large spray gentle buyer hover flock dream hybrid match whip ten mountain pitch enemy lobster afford barrel patrol desk trigger output excuse truck"
	unitInterface, err := unitMakerFunc(uuid.NewString(), mnemoPhrase)
	if err != nil {
		log.Fatal(err)
	}

	_, isCasted = unitInterface.(walletPoolUnitService)
	if !isCasted {
		log.Fatal("unable to cast pool unit Maker to named interface")
	}

	log.Printf("--- PASS: %s\n", pluginNewPoolUnitSymbol)
}
