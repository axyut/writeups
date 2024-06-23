---
title: "Challenge 11: An ECB/CBC detection oracle"
description: Set 2 of Cryptopals
---

### Problem

```plaintext
An ECB/CBC detection oracle
Now that you have ECB and CBC working:

Write a function to generate a random AES key; that's just 16 random bytes.

Write a function that encrypts data under an unknown key --- that is, a function that generates a random key and encrypts under it.

The function should look like:

encryption_oracle(your-input)
=> [MEANINGLESS JIBBER JABBER]
Under the hood, have the function append 5-10 bytes (count chosen randomly) before the plaintext and 5-10 bytes after the plaintext.

Now, have the function choose to encrypt under ECB 1/2 the time, and under CBC the other half (just use random IVs each time for CBC). Use rand(2) to decide which to use.

Detect the block cipher mode the function is using each time. You should end up with a piece of code that, pointed at a block box that might be encrypting ECB or CBC, tells you which one is happening.
```

### Code

```go
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
	ecb := IsECB(cipher, ks)

	if ecb {
		fmt.Printf("Actual: %s, Detected: ECB", actual)
	} else {
		fmt.Printf("Actual: %s, Detected: CBC", actual)
	}
}

func IsECB(cipher []byte, ks int) bool {
	for i := 0; i < len(cipher)/ks-1; i++ {
		for j := i + 1; j < len(cipher)/ks; j++ {
			a := cipher[i*ks : (i+1)*ks]
			b := cipher[j*ks : (j+1)*ks]
			if bytes.Equal(a, b) {
				return true
			}
		}
	}
	return false
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
```
