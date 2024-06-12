# bc-wallet-eth-hdwallet

## Description

Implementation of **Hierarchical Deterministic Wallet** for EVM-like blockchains.

Plugin support all EVM-like blockchains. All you need it set right values of ```NetworkChainID``` 
and ```CoinType``` variables. The plugin has 2 ways to set the desired variable values:
* The first way is through using ```SetChainID(chainID int) errors``` and ```SetHdWalletCoinType(coinType int) error``` functions.
See [#Plugin API](#plugin-api) section.
* The second method is by assembling a plugin with the specified values - see [#Build](#build) section.

HdWallet-plugin is third and last part of hd-wallet applications bundle. Also, this repo contains 
Helm-chart description for deploy full HdWallet applications bundle.

Another two parts of hdwallet-bundle is:

* [bc-wallet-common-hdwallet-controller](https://github.com/crypto-bundle/bc-wallet-common-hdwallet-controller) - 
Application for control access to wallets. Create or disable wallets, get account addresses, sign transactions.

* [bc-wallet-common-hdwallet-api](https://github.com/crypto-bundle/bc-wallet-common-hdwallet-api) - 
Storage-less application for manage in-memory HD-wallets and execute session and signature requests.

### Plugin API
Implementation of HdWallet plugin contains exported functions:
* ```NewPoolUnitfunc(walletUUID string, mnemonicDecryptedData string) (interface{}, error)```
* ```GenerateMnemonic func() (string, error)```
* ```ValidateMnemonic func(mnemonic string) bool```
* ```GetChainID() int```
* ```SetChainID(chainID int) error```
* ```GetSupportedChainIDsInfo() string```
* ```GetHdWalletCoinType() int```
* ```SetHdWalletCoinType(coinType int) error```
* ```GetSupportedCoinTypesInfo() string```
* ```GetPluginName func() string```
* ```GetPluginReleaseTag func() string```
* ```GetPluginCommitID func() string```
* ```GetPluginShortCommitID func() string```
* ```GetPluginBuildNumber func() string```
* ```GetPluginBuildDateTS func() string```

Example of usage hd-wallet pool_unit you can see in [plugin/pool_unit_test.go](plugin/pool_unit_test.go) file.
Example of plugin integration in [cmd/loader_test/main.go](cmd/loader_test/main.go) file.

### Build
Plugin support build-time variables injecting. Supported variables:
* `NetworkName` - plugin blockchain name. ethereum, polygon, bsc. You can set any value to this variable, it does not affect on plugin behavior
* `NetworkChainID` - blockchain network ID, Ethereum blockchain MainNet id = 1, You can set any value,
  but if you want to get right plugin behavior - you must select one ChainID value from EVM chain list.
  You can see available chains on [chainlist.org](https://chainlist.org/), [chainid.network](https://chainid.network/) or in another sources.
* `CoinType` - HdWallet coin type. Default value for Ethereum = 60, Ethereum Classic = 61, Binance Smart Chain = 9006, etc.
  See BIP-0044 - https://github.com/satoshilabs/slips/blob/master/slip-0044.md.
  But typically all EVM-like blockchain use Ethereum coin type  = 60.
* `ReleaseTag` - release tag in TAG.SHORT_COMMIT_ID.BUILD_NUMBER format.
* `CommitID` - latest GIT commit id.
* `ShortCommitID` - first 12 characters from CommitID.
* `BuildNumber` - ci/cd build number for BuildNumber
* `BuildDateTS` - ci/cd build date in time stamp

Build example:
```bash
RACE=-race CGO_ENABLED=1 go build -trimpath ${RACE} -installsuffix cgo -gcflags all=-N \
		-ldflags "-linkmode external -extldflags -w -s \
			-X 'main.NetworkName=${NETWORK_NAME}' \
			-X 'main.NetworkChainID=${NETWORK_CHAIN_ID}' \
			-X 'main.CoinType=${HDWALLET_COIN_TYPE}' \
			-X 'main.BuildDateTS=${BUILD_DATE_TS}' \
			-X 'main.BuildNumber=${BUILD_NUMBER}' \
			-X 'main.ReleaseTag=${RELEASE_TAG}' \
			-X 'main.CommitID=${COMMIT_ID}' \
			-X 'main.ShortCommitID=${SHORT_COMMIT_ID}'" \
		-buildmode=plugin \
		-o ./build/ethereum.so \
		./plugin
```


* `ReleaseTag` - release tag in TAG.SHORT_COMMIT_ID.BUILD_NUMBER format.
* `CommitID` - latest GIT commit id.
* `ShortCommitID` - first 12 characters from CommitID.
* `BuildNumber` - ci/cd build number for BuildNumber
* `BuildDateTS` - ci/cd build date in time stamp

## Deployment

Currently, support only kubernetes deployment flow via Helm

### Kubernetes
Application must be deployed as part of bc-wallet-<BLOCKCHAIN_NAME>-hdwallet bundle.
**_bc-wallet-ethereum-hdwallet-api_** application must be started as single container in Kubernetes Pod with shared volume.

You can see example of HELM-chart deployment application in next repositories:
* [deploy/helm/hdwallet](deploy/helm/hdwallet)
* [bc-wallet-eth-hdwallet-api/deploy/helm/hdwallet](https://github.com/crypto-bundle/bc-wallet-eth-hdwallet/tree/develop/deploy/helm/hdwallet)

## Third party libraries
Some parts of this plugin picked up from another repository - [Go HD Wallet tools](https://github.com/wemeetagain/go-hdwallet)
written by [Cayman(wemeetagain)](https://github.com/wemeetagain)

## Contributors
* Author and maintainer - [@gudron (Alex V Kotelnikov)](https://github.com/gudron)

## Licence

**bc-wallet-eth-hdwallet** is licensed under the [MIT NON-AI](./LICENSE) License.