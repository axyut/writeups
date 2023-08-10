package challenge15

import (
	"errors"
	"fmt"
)

func Challenge15() {
	goodString := []byte("ICE ICE BABY\x04\x04\x04\x04")
	getGood, err := Unpad(goodString)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Good string: %q\nUnpadded:%q", goodString, getGood)

	badString1 := []byte("ICE ICE BABY\x05\x05\x05\x05")
	badString2 := []byte("ICE ICE BABY\x01\x02\x03\x04")
	getBad1, err := Unpad(badString1)
	if err != nil {
		fmt.Println(err)
		return
	}
	getBad2, err := Unpad(badString2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Bad string: %q\nUnpadded:%q", badString1, getBad1)
	fmt.Printf("Bad string: %q\nUnpadded:%q", badString2, getBad2)

}

func Unpad(input []byte) ([]byte, error) {
	if input == nil || len(input) == 0 {
		return nil, nil
	}

	pc := input[len(input)-1]
	padLen := int(pc)

	if exists := checkPaddingIsValid(input, padLen); !exists {
		return nil, errors.New("Invalid padding, Cannot unpad.")
	}
	return input[:len(input)-padLen], nil
}
func checkPaddingIsValid(input []byte, paddingLength int) bool {
	if len(input) < paddingLength {
		return false
	}
	p := input[len(input)-(paddingLength):]
	for _, pc := range p {
		if uint(pc) != uint(len(p)) {
			return false
		}
	}
	return true
}
