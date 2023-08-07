package challenge11

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"
	c7 "cryptopals/Set1/challenge7"
	c10 "cryptopals/Set2/challenge10"
	c9 "cryptopals/Set2/challenge9"
	"encoding/hex"
	"fmt"
	"log"
)

func Challenge11() {
	ks := 16
	fill := make([]byte, ks*4)
	cipher, actual := ECBCBC_EncryptOracle(fill)

	for i := 0; i < len(cipher)/ks; i++ {
		fmt.Printf("%s\n", hex.EncodeToString(cipher[i*ks:(i+1)*ks]))
	}
	ecb := false
	for i := 0; i < len(cipher)/ks-1; i++ {
		for j := i + 1; j < len(cipher)/ks; j++ {
			a := cipher[i*ks : (i+1)*ks]
			b := cipher[j*ks : (j+1)*ks]
			if bytes.Equal(a, b) {
				ecb = true
			}
		}
	}

	if ecb {
		fmt.Printf("Actual: %s, Detected: ECB", actual)
	} else {
		fmt.Printf("Actual: %s, Detected: CBC", actual)
	}
}

func ECBCBC_EncryptOracle(input []byte) ([]byte, string) {
	prefix := RandBytes(int(RandByte()%6) + 5)
	suffix := RandBytes(int(RandByte()%6) + 5)
	res := make([]byte, len(prefix)+len(input)+len(suffix))
	actual := ""

	copy(res, prefix)
	copy(res[len(prefix):], input)
	copy(res[len(prefix)+len(input):], suffix)

	ks := 16
	key := RandBytes(ks)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}
	mode := RandByte() % 2
	res = c9.PKCS(res, ks)
	dst := make([]byte, len(res))
	if mode == 0 {
		c7.EncryptAES_ECB(block, dst, res)
		actual = "ECB"
	} else {
		IV := RandBytes(ks)
		c10.Encrypt_CBC(block, IV, dst, res)
		actual = "CBC"
	}
	return dst, actual
}

func RandBytes(n int) []byte {
	res := make([]byte, n)
	_, err := rand.Read(res)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
func RandByte() byte {
	res := make([]byte, 1)
	_, err := rand.Read(res)
	if err != nil {
		log.Fatal(err)
	}
	return res[0]
}
