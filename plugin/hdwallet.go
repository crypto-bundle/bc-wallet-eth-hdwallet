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
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"math/big"

	"github.com/btcsuite/btcd/btcutil/base58"
)

// hdWallet defines the components of a hierarchical deterministic hdwallet
type hdWallet struct {
	prvMagic    [4]byte
	pubMagic    [4]byte
	Vbytes      []byte // 4 bytes
	Depth       uint16 // 1 byte
	Fingerprint []byte // 4 bytes
	I           []byte // 4 bytes
	Chaincode   []byte // 32 bytes
	Key         []byte // 33 bytes
}

// Child returns the ith child of hdwallet w. Values of i >= 2^31
// signify private key derivation. Attempting private key derivation
// with a public key will throw an error.
func (w *hdWallet) Child(i uint32) (*hdWallet, error) {
	var fingerprint, I, newkey []byte
	switch {
	case bytes.Equal(w.Vbytes, w.prvMagic[:4]):
		pub := privToPub(w.Key)
		mac := hmac.New(sha512.New, w.Chaincode)
		if i >= uint32(0x80000000) {
			_, writeErr := mac.Write(append(w.Key, uint32ToByte(i)...))
			if writeErr != nil {
				return nil, writeErr
			}
		} else {
			_, writeErr := mac.Write(append(pub, uint32ToByte(i)...))
			if writeErr != nil {
				return nil, writeErr
			}
		}

		I = mac.Sum(nil)
		iL := new(big.Int).SetBytes(I[:32])
		if iL.Cmp(curve.N) >= 0 || iL.Sign() == 0 {
			return &hdWallet{}, errors.New("invalid child")
		}
		newkey = addPrivKeys(I[:32], w.Key)
		raw, err := hash160(privToPub(w.Key))
		if err != nil {
			return nil, err
		}

		fingerprint = raw[:4]

	case bytes.Equal(w.Vbytes, w.pubMagic[:4]):
		mac := hmac.New(sha512.New, w.Chaincode)
		if i >= uint32(0x80000000) {
			return &hdWallet{}, errors.New("can't do private derivation on public key")
		}
		_, writeErr := mac.Write(append(w.Key, uint32ToByte(i)...))
		if writeErr != nil {
			return nil, writeErr
		}

		I = mac.Sum(nil)
		iL := new(big.Int).SetBytes(I[:32])
		if iL.Cmp(curve.N) >= 0 || iL.Sign() == 0 {
			return &hdWallet{}, errors.New("invalid child")
		}

		newkey = addPubKeys(privToPub(I[:32]), w.Key)
		raw, err := hash160(w.Key)
		if err != nil {
			return nil, err
		}

		fingerprint = raw[:4]
	}
	return &hdWallet{w.prvMagic, w.pubMagic, w.Vbytes, w.Depth + 1, fingerprint, uint32ToByte(i), I[32:], newkey}, nil
}

// Serialize returns the serialized form of the hdwallet.
func (w *hdWallet) Serialize() ([]byte, error) {
	depth := uint16ToByte(w.Depth % 256)

	bindata := make([]byte, 78)
	copy(bindata, w.Vbytes)
	copy(bindata[4:], depth)
	copy(bindata[5:], w.Fingerprint)
	copy(bindata[9:], w.I)
	copy(bindata[13:], w.Chaincode)
	copy(bindata[45:], w.Key)
	raw, err := dblSha256(bindata)
	if err != nil {
		return nil, err
	}

	chksum := raw[:4]

	return append(bindata, chksum...), nil
}

// String returns the base58-encoded string form of the hdwallet.
func (w *hdWallet) String() (string, error) {
	serialized, err := w.Serialize()
	if err != nil {
		return "", err
	}

	return base58.Encode(serialized), nil
}

// hdWalletFromString returns a hdwallet given a base58-encoded extended key
func hdWalletFromString(data string, prvMagic, pubMagic [4]byte) (*hdWallet, error) {
	dbin := base58.Decode(data)
	if err := byteCheck(dbin, prvMagic, pubMagic); err != nil {
		return &hdWallet{}, err
	}

	sha256calc, err := dblSha256(dbin[:(len(dbin) - 4)])
	if err != nil {
		return &hdWallet{}, err
	}

	if !bytes.Equal(sha256calc[:4], dbin[(len(dbin)-4):]) {
		return &hdWallet{}, errors.New("invalid checksum")
	}

	vbytes := dbin[0:4]
	depth := byteToUint16(dbin[4:5])
	fingerprint := dbin[5:9]
	i := dbin[9:13]
	chaincode := dbin[13:45]
	key := dbin[45:78]

	return &hdWallet{
		prvMagic:    prvMagic,
		pubMagic:    pubMagic,
		Vbytes:      vbytes,
		Depth:       depth,
		Fingerprint: fingerprint,
		I:           i,
		Chaincode:   chaincode,
		Key:         key,
	}, nil
}

// Pub returns a new hdwallet which is the public key version of w.
// If w is a public key, Pub returns a copy of w
func (w *hdWallet) Pub() *hdWallet {
	if bytes.Equal(w.Vbytes, w.pubMagic[:4]) {
		return &hdWallet{w.prvMagic, w.pubMagic, w.Vbytes, w.Depth, w.Fingerprint, w.I, w.Chaincode, w.Key}
	}

	return &hdWallet{w.prvMagic, w.pubMagic, w.pubMagic[:4], w.Depth, w.Fingerprint, w.I, w.Chaincode, privToPub(w.Key)}
}

// stringChild returns the ith base58-encoded extended key of a base58-encoded extended key.
func stringChild(data string, i uint32, prvMagic, pubMagic [4]byte) (string, error) {
	w, err := hdWalletFromString(data, prvMagic, pubMagic)
	if err != nil {
		return "", err
	}

	w, err = w.Child(i)
	if err != nil {
		return "", err
	}

	str, err := w.String()
	if err != nil {
		return "", err
	}

	return str, nil
}

// stringAddress returns the Bitcoin address of a base58-encoded extended key.
func stringAddress(data string, prvMagic, pubMagic [4]byte) (string, error) {
	w, err := hdWalletFromString(data, prvMagic, pubMagic)
	if err != nil {
		return "", err
	}

	addr, err := w.Address()
	if err != nil {
		return "", err
	}

	return addr, nil
}

// Address returns bitcoin address represented by hd-wallet w.
func (w *hdWallet) Address() (string, error) {
	x, y := expand(w.Key)
	paddedKey, err := hex.DecodeString("04")
	if err != nil {
		return "", err
	}

	paddedKey = append(paddedKey, append(x.Bytes(), y.Bytes()...)...)
	addr1, err := hex.DecodeString("00")
	if err != nil {
		return "", err
	}

	raw, err := hash160(paddedKey)
	if err != nil {
		return "", err
	}

	addr1 = append(addr1, raw...)
	chkSum, err := dblSha256(addr1)
	if err != nil {
		return "", err
	}

	return base58.Encode(append(addr1, chkSum[:4]...)), nil
}

// genSeed returns a random seed with a length measured in bytes.
// The length must be at least 128.
func genSeed(length int) ([]byte, error) {
	b := make([]byte, length)
	if length < 128 {
		return b, errors.New("length must be at least 128 bits")
	}
	_, err := rand.Read(b)
	return b, err
}

// masterKey returns a new hdwallet given a random seed.
func masterKey(seed []byte, prvMagic, pubMagic [4]byte) (*hdWallet, error) {
	key := []byte("Bitcoin seed")
	mac := hmac.New(sha512.New, key)
	_, err := mac.Write(seed)
	if err != nil {
		return nil, err
	}

	I := mac.Sum(nil)
	secret := I[:len(I)/2]
	chainCode := I[len(I)/2:]
	depth := 0
	i := make([]byte, 4)
	fingerprint := make([]byte, 4)
	zero := make([]byte, 1)

	return &hdWallet{
		prvMagic:    prvMagic,
		pubMagic:    pubMagic,
		Vbytes:      prvMagic[:4],
		Depth:       uint16(depth),
		Fingerprint: fingerprint,
		I:           i,
		Chaincode:   chainCode,
		Key:         append(zero, secret...),
	}, nil
}

func byteCheck(dBin []byte, prvMagic, pubMagic [4]byte) error {
	// check proper length
	if len(dBin) != 82 {
		return errors.New("invalid string")
	}
	// check for correct Public or Private vbytes
	if !bytes.Equal(dBin[:4], pubMagic[:4]) && !bytes.Equal(dBin[:4], prvMagic[:4]) {
		return errors.New("invalid string")
	}

	// if Public, check x coord is on curve
	x, y := expand(dBin[45:78])
	if bytes.Equal(dBin[:4], pubMagic[:4]) {
		if !onCurve(x, y) {
			return errors.New("invalid string")
		}
	}
	return nil
}
