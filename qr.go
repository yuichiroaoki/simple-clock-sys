package main

import (
	"bytes"
	"fmt"
	"image/png"
	"os"
	"os/exec"
	"strings"

	"github.com/skip2/go-qrcode"
)

// Create a QR code from a hash and save it to a file
func generateQrCodeFromHash(hash, path string) (*qrcode.QRCode, error) {
	// Create the QR code
	qrCode, err := qrcode.New(hash, qrcode.Medium)
	if err != nil {
		fmt.Println("Error creating QR code:", err)
		return nil, err
	}

	// Save the QR code to a file
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return nil, err
	}
	defer file.Close()

	err = png.Encode(file, qrCode.Image(256))
	if err != nil {
		fmt.Println("Error encoding QR code as PNG:", err)
		return nil, err
	}

	return qrCode, nil
}

// Get the hash from a QR code
func getHashFromQrCode(q qrcode.QRCode) (string, error) {
	hash, err := zbarimgDecode(q)
	if err != nil {
		fmt.Println("Error decoding QR code:", err)
		return "", err
	}

	return hash, nil
}

func zbarimgDecode(q qrcode.QRCode) (string, error) {
	var png []byte

	png, err := q.PNG(256)
	if err != nil {
		return "", err
	}

	cmd := exec.Command("zbarimg", "--quiet", "-Sdisable",
		"-Sqrcode.enable", "-")

	var out bytes.Buffer

	cmd.Stdin = bytes.NewBuffer(png)
	cmd.Stdout = &out

	err = cmd.Run()

	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(strings.TrimPrefix(out.String(), "QR-Code:"), "\n"), nil
}
