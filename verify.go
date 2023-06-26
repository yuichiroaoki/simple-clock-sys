package main

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// ref. https://gist.github.com/dcb9/385631846097e1f59e3cba3b1d42f3ed#file-eth_sign_verify-go
func verifySig(from, sigHex string, msg []byte) bool {
	sig := hexutil.MustDecode(sigHex)

	msg = accounts.TextHash(msg)
	if sig[crypto.RecoveryIDOffset] == 27 || sig[crypto.RecoveryIDOffset] == 28 {
		sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}

	recovered, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*recovered)

	checksummedFrom := hashToAddress(from)

	return checksummedFrom == recoveredAddr.Hex()
}

func hashToAddress(hash string) string {
	// Remove the "0x" prefix if present
	hash = strings.TrimPrefix(hash, "0x")

	// Convert the hash string to bytes
	hashBytes := common.FromHex(hash)

	// Create an Ethereum address from the hash bytes
	address := common.BytesToAddress(hashBytes)

	// Generate the checksummed address string
	checksummedAddress := address.Hex()
	return checksummedAddress
}
