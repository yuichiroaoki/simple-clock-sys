package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// clocking system device

	// create a random message
	data := []byte("random message")
	hash := accounts.TextHash(data)

	// create a QR code from the hash
	// show it to the client
	// get the hash from the QR code

	// client mobile app
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}
	address := privateKeyToAddress(privateKey)

	// sign the message from the clocking system device
	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	sigHash := hexutil.Encode(signature)

	fmt.Println("Signature:    ", sigHash)

	path := "images/qrcode.png"
	q, err := generateQrCodeFromHash(sigHash, path)
	if err != nil {
		fmt.Println("Error generating QR code:", err)
		return
	}

	// clocking system device
	qrcodeHash, err := getHashFromQrCode(*q)
	if err != nil {
		fmt.Println("Error getting hash from QR code:", err)
		return
	}

	fmt.Println("QR code hash: ", qrcodeHash)
	fmt.Println("Verified: ", verifySig(address, qrcodeHash, data))
	// start clocking
}
