---
title: "Challenge 12: Byte-at-a-time ECB decryption (Simple)"
description: Set 2 of Cryptopals
---

### Problem

```plaintext
Byte-at-a-time ECB decryption (Simple)
Copy your oracle function to a new function that encrypts buffers under ECB mode using a consistent but unknown key (for instance, assign a single random key, once, to a global variable).

Now take that same function and have it append to the plaintext, BEFORE ENCRYPTING, the following string:

Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg
aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq
dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg
YnkK
Spoiler alert.
Do not decode this string now. Don't do it.

Base64 decode the string before appending it. Do not base64 decode the string by hand; make your code do it. The point is that you don't know its contents.

What you have now is a function that produces:

AES-128-ECB(your-string || unknown-string, random-key)
It turns out: you can decrypt "unknown-string" with repeated calls to the oracle function!

Here's roughly how:

Feed identical bytes of your-string to the function 1 at a time --- start with 1 byte ("A"), then "AA", then "AAA" and so on. Discover the block size of the cipher. You know it, but do this step anyway.
Detect that the function is using ECB. You already know, but do this step anyways.
Knowing the block size, craft an input block that is exactly 1 byte short (for instance, if the block size is 8 bytes, make "AAAAAAA"). Think about what the oracle function is going to put in that last byte position.
Make a dictionary of every possible last byte by feeding different strings to the oracle; for instance, "AAAAAAAA", "AAAAAAAB", "AAAAAAAC", remembering the first block of each invocation.
Match the output of the one-byte-short input to one of the entries in your dictionary. You've now discovered the first byte of unknown-string.
Repeat for the next byte.
```

### Code

```go
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
```
