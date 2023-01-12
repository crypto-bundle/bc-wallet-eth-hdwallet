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

package hdwallet

import (
	"fmt"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

// NewBTCWaller create new wallet
func (w *Wallet) NewBtcWallet(account, change, address uint32) (*BTC, error) {
	blockChainParams := chaincfg.MainNetParams

	blockChainParams.HDPrivateKeyID = [4]byte{0x04, 0x9d, 0x78, 0x78} // yprv
	blockChainParams.HDPublicKeyID = [4]byte{0x04, 0x9d, 0x7c, 0xb2}  // ypub

	accountKey, key, err := w.GetChildKey(&blockChainParams, DefaultPurpose, BtcCoinNumber, account, change, address)
	if err != nil {
		return nil, err
	}

	return &BTC{
		blockChainParams: &blockChainParams,
		purpose:          DefaultPurpose,
		coinType:         BtcCoinNumber,
		account:          account,
		change:           change,
		addressNumber:    address,

		ExtendedKey: key,
		AccountKey:  accountKey,
	}, nil
}

// GetAddress get address string
func (b *BTC) GetAddress() (string, error) {
	return b.ExtendedKey.AddressP2WPKHInP2SH()
}

// GetP2WPKHAddress get address string
func (b *BTC) GetP2WPKHAddress() (string, error) {
	return b.ExtendedKey.AddressP2WPKH()
}

// GetP2WPKHAddress get address string
func (b *BTC) GetP2SHAddress() (string, error) {
	return b.ExtendedKey.AddressP2WPKHInP2SH()
}

// GetP2WPKHAddress get address string
func (b *BTC) GetP2PKHAddress() (string, error) {
	return b.ExtendedKey.AddressP2PKH()
}

// GetPrvKey get address private key
func (b *BTC) GetPrvKey() (string, error) {
	prvKey, err := btcutil.NewWIF(b.ExtendedKey.Private, b.ExtendedKey.Network, true)
	if err != nil {
		return "", nil
	}
	return prvKey.String(), nil
}

// GetPrvKey get address private key
func (b *BTC) GetWIF() (*btcutil.WIF, error) {
	prvKey, err := btcutil.NewWIF(b.ExtendedKey.Private, b.ExtendedKey.Network, true)
	if err != nil {
		return nil, err
	}

	return prvKey, nil
}

// GetPubKey get address public key
func (b *BTC) GetPubKey() string {
	return b.ExtendedKey.PublicHex()
}

// AccountPrvKey return string key
func (b *BTC) AccountPrvKey() string {
	return b.AccountKey.Private
}

// AccountPrvKeyNoMagic return string key
func (b *BTC) AccountPrvKeyNoMagic() string {
	return b.AccountKey.Private[4:]
}

// AccountPubKey return string key
func (b *BTC) AccountPubKey() string {
	return b.AccountKey.Public
}

// AccountPubKey return string key
func (b *BTC) AccountPubKeyNoMagic() string {
	return b.AccountKey.Public[4:]
}

// AccountPubKey return string key
func (b *BTC) PKH() (string, error) {
	return b.ExtendedKey.AddressP2WPKH()
}

// AccountPubKey return string key
func (b *BTC) HEX() string {
	return b.ExtendedKey.PublicHex()
}

// AccountPubKey return string key
func (b *BTC) HASH() ([]byte, error) {
	return b.ExtendedKey.PublicHash()
}

// GetPrvKey get address private key
func (b *BTC) GetAccountWIF() (*btcutil.WIF, error) {
	ecKey, err := b.AccountKey.ExtendedKey.ECPrivKey()
	if err != nil {
		return nil, err
	}

	wif, err := btcutil.NewWIF(ecKey, b.ExtendedKey.Network, true)
	if err != nil {
		return nil, err
	}

	return wif, nil
}

// GetPurpose
func (b *BTC) GetPurpose() int {
	return b.purpose
}

// GetBlockChainConfig
func (b *BTC) GetBlockChainConfig() *chaincfg.Params {
	return b.blockChainParams
}

// GetPath
func (b *BTC) GetPath() string {
	return fmt.Sprintf("m/%d'/%d'/%d'/%d/%d",
		b.GetPurpose(), b.GetCoinType(), b.account, b.change, b.addressNumber)
}

// GetCoinNumber
func (b *BTC) GetCoinType() int {
	return b.coinType
}
