---
title: "Challenge 16: CBC bitflipping attacks"
description: Set 2 of Cryptopals
---

### Problem

```plaintext
CBC bitflipping attacks
Generate a random AES key.

Combine your padding code and CBC code to write two functions.

The first function should take an arbitrary input string, prepend the string:

"comment1=cooking%20MCs;userdata="
.. and append the string:

";comment2=%20like%20a%20pound%20of%20bacon"
The function should quote out the ";" and "=" characters.

The function should then pad out the input to the 16-byte AES block length and encrypt it under the random AES key.

The second function should decrypt the string and look for the characters ";admin=true;" (or, equivalently, decrypt, split the string on ";", convert each resulting string into 2-tuples, and look for the "admin" tuple).

Return true or false based on whether the string exists.

If you've written the first function properly, it should not be possible to provide user input to it that will generate the string the second function is looking for. We'll have to break the crypto to do that.

Instead, modify the ciphertext (without knowledge of the AES key) to accomplish this.

You're relying on the fact that in CBC mode, a 1-bit error in a ciphertext block:

Completely scrambles the block the error occurs in
Produces the identical 1-bit error(/edit) in the next ciphertext block.
```

### Code

```go
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
```
