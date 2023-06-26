package main

import (
	"crypto/ecdsa"
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestVerifySig(t *testing.T) {
	// https://etherscan.io/verifySig/20737
	data := []byte("Hey Guys! Never doing that again!")
	isValid := verifySig(
		"0x18d4d9e749e8b27e00d775ebe442bbaf67920e0e",
		"0xfdbecbe0683c544bb4868d78c46624b35dd41229ac8de5382c23a517deb1a9024c581da3acdfb4c6bd2720230a084c039d77ed0627235bc739631a2e2c30ee741c",
		data,
	)
	if !isValid {
		t.Error("verifySig should return true")
	}

	// https://etherscan.io/verifySig/20738
	data = []byte(`IS THIS MY CONTRACT ADRESS ? 
I DON'T UNDERSTAND WHAT TO DO TO GET MY MONEY !!`)
	isValid = verifySig(
		"0x885C0044D2EeB1eDc7E08C52755E0867bd033A65",
		"0x5dd790e3c42965862d3e360723a654a88c65ef6384a6d5dcb12416edcb68618b4a18c7ec3c1242541d64d6b3a42c2f4d722619d55fb279d0cf588cbbbcd463201b",
		data,
	)
	if !isValid {
		t.Error("verifySig should return true")
	}

	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}
	messageHash := accounts.TextHash(data)
	signature, err := crypto.Sign(messageHash, privateKey)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	from := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	sigHash := hexutil.Encode(signature)
	isValid = verifySig(
		from,
		sigHash,
		data,
	)
	if !isValid {
		t.Error("verifySig should return true")
	}

}

func TestHashToAddress(t *testing.T) {
	var tests = []struct {
		Input  string
		Output string
	}{
		// Test cases from https://github.com/ethereum/EIPs/blob/master/EIPS/eip-55.md#specification
		{"0x5aaeb6053f3e94c9b9a09f33669435e7ef1beaed", "0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed"},
		{"0xfb6916095ca1df60bb79ce92ce3ea74c37c5d359", "0xfB6916095ca1df60bB79Ce92cE3Ea74c37c5d359"},
		{"0xdbf03b407c01e7cd3cbea99509d93f8dddc8c6fb", "0xdbF03B407c01E7cD3CBea99509d93f8DDDC8C6FB"},
		{"0xd1220a0cf47c7b9be7a2e6ba89f429762e7b9adb", "0xD1220A0cf47c7B9Be7A2E6BA89F429762e7b9aDb"},
	}

	for _, test := range tests {
		result := hashToAddress(test.Input)
		if result != test.Output {
			t.Errorf("hashToAddress(%q) = %q, want %q", test.Input, result, test.Output)
		}
	}
}
