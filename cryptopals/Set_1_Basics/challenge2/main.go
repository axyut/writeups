package main

import (
	"encoding/hex"
	"errors"
	"fmt"
)

/*
Write a function that takes two equal-length buffers and produces their XOR combination.
If your function works properly, then when you feed it the string:

1c0111001f010100061a024b53535009181c
... after hex decoding, and when XOR'd against:
686974207468652062756c6c277320657965
... should produce:
746865206b696420646f6e277420706c6179
*/

func main() {
	input1 := []byte("1c0111001f010100061a024b53535009181c")
	input2 := []byte("686974207468652062756c6c277320657965")
	output, err := xorEncrypt(input1, input2)

	if err != nil {
		fmt.Printf("Error Occured %s", err)
		return
	}
	fmt.Printf("%s\n", output)
}

// Takes raw bytes
func xorEncrypt(in, mid []byte) ([]byte, error) {
	input1 := HexaDecode(in)
	input2 := HexaDecode(mid)

	if len(input1) != len(input2) {
		return nil, errors.New("Provide Equal Length buffers")
	}
	out := make([]byte, len(input1))
	for i := 0; i < len(input1); i++ {
		out[i] = input1[i] ^ input2[i]
	}
	ret := HexaEncode(out)
	return ret, nil
}

// Decoding form Hex
func HexaDecode(hexEncoded []byte) []byte {
	decoded := make([]byte, hex.DecodedLen(len(hexEncoded)))

	_, err := hex.Decode(decoded, hexEncoded)
	if err != nil {
		fmt.Printf("Failed to decode hex: %s", err)
		return nil
	}
	return decoded
}

// Encoding to Hex
func HexaEncode(in []byte) []byte {
	encoded := make([]byte, hex.EncodedLen(len(in)))
	hex.Encode(encoded, in)
	return encoded
}
