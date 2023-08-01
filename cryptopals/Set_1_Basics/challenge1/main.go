package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	// HEX Decode
	hexEncoded := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	decoded := make([]byte, len(hexEncoded))

	_, err := hex.Decode(decoded, hexEncoded)
	if err != nil {
		fmt.Printf("Failed to decode hex: %s", err)
		return
	}
	//fmt.Println(len)
	//fmt.Println(hex.DecodedLen(len(hexEncoded)))
	fmt.Printf("Hex Decoded: %s", decoded)

	// Base64 Encode
	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(decoded)))
	base64.StdEncoding.Encode(encoded, decoded)
	fmt.Printf("\nBase64: %s", encoded)
}
