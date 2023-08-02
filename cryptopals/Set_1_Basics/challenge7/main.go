package main

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"os"
	//"strings"
)

func main() {
	key := []byte("YELLOW SUBMARINE")
	file, err := os.ReadFile("7.txt")
	//fmt.Printf("%s", file)
	if err != nil {
		return
	}
	//lines := strings.Split(string(file), "\n")
	bitByte := []byte(file)
	decoded := DecodeOne(bitByte)

	ecb, err := decryptAESECB(decoded, key, 16)
	if err != nil {
		return
	}
	fmt.Printf("%s", ecb)
}

func decryptAESECB(cipherText []byte, key []byte, blockSize int) ([]byte, error) {
	plainText := make([]byte, len(cipherText))
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("could not initialize AES: %w", err)
	}

	for i := 0; i < (len(plainText) / blockSize); i++ {
		start := i * blockSize
		end := (i + 1) * blockSize
		cipher.Decrypt(plainText[start:end], cipherText[start:end])
	}

	return plainText, nil
}

func DecodeOne(input []byte) []byte {
	decoded := make([]byte, base64.RawStdEncoding.DecodedLen(len(input)))
	_, err := base64.RawStdEncoding.Decode(decoded, input)
	if err != nil {
		fmt.Printf("%s", err)
		return nil
	}
	return decoded
}
