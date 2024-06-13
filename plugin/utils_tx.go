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
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var big8 = big.NewInt(8)

func recoverPlain(sighash common.Hash,
	R, S, Vb *big.Int,
	homestead bool,
) (*ecdsa.PublicKey, common.Address, error) {
	if Vb.BitLen() > 8 {
		return nil, common.Address{}, errors.New("invalidSign")
	}

	V := byte(Vb.Uint64() - 27)
	if !crypto.ValidateSignatureValues(V, R, S, homestead) {
		return nil, common.Address{}, errors.New("invalidSign")
	}

	// encode the signature in uncompressed format
	r, s := R.Bytes(), S.Bytes()
	sig := make([]byte, crypto.SignatureLength)
	copy(sig[32-len(r):32], r)
	copy(sig[64-len(s):64], s)
	sig[64] = V

	// DIRTY HACK
	// https://stackoverflow.com/questions/49085737/geth-ecrecover-invalid-signature-recovery-id
	// https://gist.github.com/dcb9/385631846097e1f59e3cba3b1d42f3ed#file-eth_sign_verify-go
	if sig[crypto.RecoveryIDOffset] == 27 || sig[crypto.RecoveryIDOffset] == 28 {
		sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}

	// recover the public key from the signature
	pub, err := crypto.Ecrecover(sighash[:], sig)
	if err != nil {
		return nil, common.Address{}, err
	}

	if len(pub) == 0 || pub[0] != 4 {
		return nil, common.Address{}, errors.New("invalid public key")
	}

	var addr common.Address
	copy(addr[:], crypto.Keccak256(pub[1:])[12:])

	ECDSAPub, err := crypto.UnmarshalPubkey(pub)
	if err != nil {
		return nil, common.Address{}, err
	}

	return ECDSAPub, addr, nil
}

func extractECSDAPublicKey(tx *types.Transaction) (*ecdsa.PublicKey, common.Address, error) {
	V, R, S := tx.RawSignatureValues()
	var signer types.Signer = nil

	switch tx.Type() {
	case types.LegacyTxType:
		chaiIDMul := new(big.Int).Mul(tx.ChainId(), big.NewInt(2))
		v := new(big.Int).Sub(V, chaiIDMul)
		V = v.Sub(v, big8)

		signer = types.NewEIP155Signer(tx.ChainId())

	case types.AccessListTxType:
		// AL txs are defined to use 0 and 1 as their recovery
		// id, add 27 to become equivalent to unprotected Homestead signatures.
		V = new(big.Int).Add(V, big.NewInt(27))

		signer = types.NewEIP2930Signer(tx.ChainId())

	case types.DynamicFeeTxType:
		// DynamicFee txs are defined to use 0 and 1 as their recovery
		// id, add 27 to become equivalent to unprotected Homestead signatures.
		V = new(big.Int).Add(V, big.NewInt(27))

		signer = types.NewLondonSigner(tx.ChainId())

	case types.BlobTxType:
		// Blob txs are defined to use 0 and 1 as their recovery
		// id, add 27 to become equivalent to unprotected Homestead signatures.
		V = new(big.Int).Add(V, big.NewInt(27))

		signer = types.NewCancunSigner(tx.ChainId())
	default:
		// Blob txs are defined to use 0 and 1 as their recovery
		// id, add 27 to become equivalent to unprotected Homestead signatures.
		V = new(big.Int).Add(V, big.NewInt(27))

		signer = types.LatestSignerForChainID(tx.ChainId())
	}

	return recoverPlain(signer.Hash(tx), R, S, V, true)
}
