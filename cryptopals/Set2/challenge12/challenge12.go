package challenge12

import (
	"bytes"
	"crypto/aes"
	c6 "cryptopals/Set1/challenge6"
	c7 "cryptopals/Set1/challenge7"
	c11 "cryptopals/Set2/challenge11"
	c9 "cryptopals/Set2/challenge9"

	"fmt"
	"log"
)

var key []byte
var decoded []byte

func init() {
	key = c11.RandBytes(16)
	var err error

	decoded, err = c6.Decodebase64LineByLine("./Set2/challenge12/12.txt")
	if err != nil {
		panic("cannot read uknown from a file")
	}
}

func Challenge12() {
	size := blockSizeGen()

	cipher := Encrypt_Oracle(make([]byte, size*4))
	if !c11.IsECB(cipher, size) {
		panic("Not ECB")
	}
	fmt.Println("ECB detected.")

	cipher = Encrypt_Oracle(make([]byte, 0))
	max := len(cipher)
	fmt.Println("Length of decoded msg: ", max)

	fill := make([]byte, size)
	cracked := make([]byte, max)

	for blockLen := 0; blockLen < max/size; blockLen++ {

		for i := 1; i < 17 && blockLen*size+i-1 < max; i++ {

			pos := blockLen*size + i - 1

			if pos >= max {
				break
			}
			if blockLen == 0 {
				copy(fill[len(fill)-i:], cracked[:i])
			} else {
				copy(fill, cracked[pos-size+1:pos])
			}

			cipher := Encrypt_Oracle(fill[:size-i])
			ideal := cipher[blockLen*size : (blockLen+1)*size]

			found := false

			for block := 0; block < 256; block++ {
				fill[size-1] = byte(block)
				cipher := Encrypt_Oracle(fill)

				if bytes.Equal(ideal, cipher[:size]) {
					cracked[pos] = byte(block)
					found = true
					break
				}
			}

			if !found {
				if len(cracked)-pos < 16 {
					fmt.Printf("Decrypted:\n%q\n", cracked)
					break
				}
				fmt.Printf("Decrypted:\n%q\n", cracked)
				break
			}
		}
	}
}

func Encrypt_Oracle(input []byte) []byte {
	src := append(input, decoded...)
	ks := len(key)

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}

	src = c9.PKCS(src, ks)

	dst := make([]byte, len(src))
	c7.EncryptAES_ECB(block, dst, src)

	return dst
}

func blockSizeGen() int {
	size := 0
	prevLen := len(Encrypt_Oracle(make([]byte, 1)))

	for i := 2; i < 32; i++ {
		cipher := Encrypt_Oracle(make([]byte, i))
		if prevLen != len(cipher) {
			size = len(cipher) - prevLen
			break
		}
	}
	return size
}
