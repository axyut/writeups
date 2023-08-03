package challenge5

import (
	"fmt"

	c1 "github.com/axyut/cryptopals/challenge1"
)

func Challenge5() {
	plaintext := []byte(`Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`)

	value := EncryptWithRepeatingKey(plaintext, []byte("ICE"))

	encoded := c1.EncodeHex(value)

	fmt.Printf("%s", encoded)
}

func EncryptWithRepeatingKey(payload, key []byte) []byte {
	cipher := make([]byte, len(payload))
	for i := 0; i < len(payload); i++ {
		cipher[i] = payload[i] ^ key[i%len(key)]
	}
	return cipher
}
