package challenge2

import (
	"errors"
	"fmt"

	c1 "github.com/axyut/cryptopals/challenge1"
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

func Challenge2() {
	input1 := []byte("1c0111001f010100061a024b53535009181c")
	input2 := []byte("686974207468652062756c6c277320657965")
	output, err := XorEncrypt(input1, input2)

	if err != nil {
		fmt.Printf("Error Occured %s", err)
		return
	}
	fmt.Printf("%s\n", output)
}

// Takes raw bytes
func XorEncrypt(in, mid []byte) ([]byte, error) {
	input1 := c1.DecodeHex(in)
	input2 := c1.DecodeHex(mid)

	if len(input1) != len(input2) {
		return nil, errors.New("Provide Equal Length buffers")
	}
	out := make([]byte, len(input1))
	for i := 0; i < len(input1); i++ {
		out[i] = input1[i] ^ input2[i]
	}
	ret := c1.EncodeHex(out)
	return ret, nil
}
