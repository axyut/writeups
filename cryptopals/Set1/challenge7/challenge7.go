package challenge7

import (
	"crypto/aes"
	"fmt"
	"os"

	//"strings"
	c1 "cryptopals/Set1/challenge1"
)

func Challenge7() {
	key := []byte("YELLOW SUBMARINE")
	file, err := os.ReadFile("./challenge7/7.txt")
	//fmt.Printf("%s", file)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	//lines := strings.Split(string(file), "\n")
	bitByte := []byte(file)
	decoded := c1.DecodeBase64(bitByte)

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
