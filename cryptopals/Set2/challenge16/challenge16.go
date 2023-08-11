package challenge16

import (
	"crypto/aes"
	c10 "cryptopals/Set2/challenge10"
	c11 "cryptopals/Set2/challenge11"
	c9 "cryptopals/Set2/challenge9"
	"fmt"
	"log"
	"strings"
)

var key []byte

const prefix = "comment1=cooking%20MCs;userdata="
const suffix = ";comment2=%20like%20a%20pound%20of%20bacon"

func init() {
	key = c11.RandBytes(16)
}

func Challenge16() {
	cipher, iv := EncryptOracle([]byte("test0000000000003admin=true"))

	forge := byte(';') ^ byte('3')
	cipher[32] ^= forge

	isAdmin := DecryptOracle(cipher, iv)
	fmt.Printf("isAdmin = %+v\n", isAdmin)
}

func EncryptOracle(input []byte) ([]byte, []byte) {

	src := append([]byte(prefix), input...)
	src = append(src, []byte(suffix)...)

	keySize := len(key)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}
	src = c9.PKCS(src, keySize)
	dst := make([]byte, len(src))
	iv := c11.RandBytes(16)
	c10.Encrypt_CBC(block, iv, dst, src)
	return dst, iv
}

const adminCheck = ";admin=true;"

func DecryptOracle(cipher, iv []byte) bool {

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}

	dst := make([]byte, len(cipher))
	c10.Decrypt_CBC(block, iv, dst, cipher)

	pt, err := c9.UnpadPKCS(dst)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Contains(string(pt), adminCheck)
}
