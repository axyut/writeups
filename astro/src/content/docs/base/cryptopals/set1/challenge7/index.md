---
title: "Challenge 7: AES in ECB mode"
description: Set 1 of Cryptopals
---

### Problem

```plaintext
AES in ECB mode
The Base64-encoded content in this file has been encrypted via AES-128 in ECB mode under the key

"YELLOW SUBMARINE".
(case-sensitive, without the quotes; exactly 16 characters; I like "YELLOW SUBMARINE" because it's exactly 16 bytes long, and now you do too).

Decrypt it. You know the key, after all.

Easiest way: use OpenSSL::Cipher and give it AES-128-ECB as the cipher.
```

### Solution

-   Decode the base64 encoded string from the file
-   Use the key "YELLOW SUBMARINE" to decrypt the string
-   Write function to decrypt AES-128-ECB using key and ciphertext and return plaintext

![cyberchef](./cyberchef.png)

We can see the text file cipher and the key given to us in the problem are correct and our code is correct because the output is the same as the output from the cyberchef recipe.

### Decoded

![Decoded](./decoded.png)

### Code

```go
package challenge7

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"

	//"strings"
	c1 "cryptopals/Set1/challenge1"
)

func Challenge7() {
	key := []byte("YELLOW SUBMARINE")
	file, err := os.ReadFile("./Set1/challenge7/7.txt")
	//fmt.Printf("%s", file)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	//lines := strings.Split(string(file), "\n")
	bitByte := []byte(file)
	decoded := c1.DecodeBase64(bitByte)

	ecb, err := DecryptAES_ECB(decoded, key, 16)
	if err != nil {
		return
	}
	fmt.Printf("%s", ecb)
}

func DecryptAES_ECB(cipherText []byte, key []byte, blockSize int) ([]byte, error) {
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

// ECBDecrypt assumes that dst and src of the same length
func EncryptAES_ECB(block cipher.Block, dst, src []byte) {
	if len(src) != len(dst) {
		panic("src and dst lengths do not match")
	}
	sz := block.BlockSize()
	for i := 0; i < len(src)/sz; i++ {
		from := i * sz
		to := (i + 1) * sz
		block.Encrypt(dst[from:to], src[from:to])
	}
}
```
