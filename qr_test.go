package main

import (
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestGetHashFromQrCode(t *testing.T) {
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	data := []byte("test sign message")
	hash := crypto.Keccak256Hash(data)

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	sigHash := hexutil.Encode(signature)

	path := "images/qrcode.png"
	q, err := generateQrCodeFromHash(sigHash, path)
	if err != nil {
		t.Error("Error generating QR code:", err)
	}

	qrcodeHash, err := getHashFromQrCode(*q)
	if err != nil {
		t.Error("Error getting hash from QR code:", err)
	}

	if qrcodeHash != sigHash {
		t.Errorf("QR code hash %s does not match hash %s", qrcodeHash, sigHash)
	}
}
