package challenge9

import "fmt"

func Challenge9() {
	input := []byte("YELLOW SUBMARINE")
	padded := PKCS(input, 20)
	fmt.Printf("No padding: %q\n", input)
	fmt.Printf("Padding: %q\n", padded)
}
func PKCS(input []byte, blockSize int) []byte {
	r := len(input) % blockSize
	pl := blockSize - r

	for i := 0; i < pl; i++ {
		input = append(input, byte(pl))
	}
	return input
}
