# bc-wallet-eth-hdwallet

## Description

HdWallet-plugin is third and last part of hd-wallet applications bundle. This repository contains implementation of
**Hierarchical Deterministic Wallet** for Ethereum-like blockchain. Also, this repo contains Helm-chart description for deploy full
HdWallet applications bundle for Ethereum.  

Another two parts of hdwallet-bundle is:

* [bc-wallet-common-hdwallet-controller](https://github.com/crypto-bundle/bc-wallet-common-hdwallet-controller) - 
Application for control access to wallets. Create or disable wallets, get account addresses, sign transactions.

* [bc-wallet-common-hdwallet-api](https://github.com/crypto-bundle/bc-wallet-common-hdwallet-api) - 
Storage-less application for manage in-memory HD-wallets and execute session and signature requests.

### Ethereum And Ethereum-like HdWallet plugin
Implementation of HdWallet plugin contains exported functions:
* ```NewPoolUnitfunc(walletUUID string, mnemonicDecryptedData string) (interface{}, error)```
* ```GenerateMnemonic func() (string, error)```
* ```ValidateMnemonic func(mnemonic string) bool```
* ```GetChainID() int```
* ```GetSupportedChainIDs() []int```
* ```GetPluginName func() string```
* ```GetPluginReleaseTag func() string```
* ```GetPluginCommitID func() string```
* ```GetPluginShortCommitID func() string```
* ```GetPluginBuildNumber func() string```
* ```GetPluginBuildDateTS func() string```

Example of usage hd-wallet pool_unit you can see in [plugin/pool_unit_test.go](plugin/pool_unit_test.go) file.
Example of plugin integration in [cmd/loader_test/main.go](cmd/loader_test/main.go) file.

## Deployment

Currently, support only kubernetes deployment flow via Helm

### Kubernetes
Application must be deployed as part of bc-wallet-<BLOCKCHAIN_NAME>-hdwallet bundle.
bc-wallet-ethereum-hdwallet-api application must be started as single container in Kubernetes Pod with shared volume.

You can see example of HELM-chart deployment application in next repositories:
* [deploy/helm/hdwallet](deploy/helm/hdwallet)
* [bc-wallet-eth-hdwallet-api/deploy/helm/hdwallet](https://github.com/crypto-bundle/bc-wallet-eth-hdwallet/tree/develop/deploy/helm/hdwallet)

## Third party libraries
Some parts of this plugin picked up from another repository - [Go HD Wallet tools](https://github.com/wemeetagain/go-hdwallet)
written by [Cayman(wemeetagain)](https://github.com/wemeetagain)

## Licence

**bc-wallet-eth-hdwallet** is licensed under the [MIT NON-AI](./LICENSE) License.