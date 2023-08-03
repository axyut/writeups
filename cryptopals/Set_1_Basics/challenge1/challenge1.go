package challenge1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func Challenge1() {
	// HEX Decode
	hexEncoded := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")

	hexDecoded := DecodeHex(hexEncoded)
	fmt.Printf("Hex Decoded: %s", hexDecoded)

	// Base64 Encode
	encoded := EncodeBase64(hexDecoded)
	fmt.Printf("\nBase64: %s", encoded)
}

func DecodeHex(input []byte) []byte {
	decoded := make([]byte, hex.DecodedLen(len(input)))

	_, err := hex.Decode(decoded, input)
	if err != nil {
		fmt.Printf("Failed to decode hex: %s", err)
		return nil
	}
	return decoded
}

func EncodeHex(in []byte) []byte {
	encoded := make([]byte, hex.EncodedLen(len(in)))
	hex.Encode(encoded, in)
	return encoded
}

func EncodeBase64(input []byte) []byte {
	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(input)))
	base64.StdEncoding.Encode(encoded, input)
	return encoded
}

func DecodeBase64(input []byte) []byte {
	decoded := make([]byte, base64.RawStdEncoding.DecodedLen(len(input)))
	_, err := base64.RawStdEncoding.Decode(decoded, input)
	if err != nil {
		fmt.Printf("%s", err)
		return nil
	}
	return decoded
}
