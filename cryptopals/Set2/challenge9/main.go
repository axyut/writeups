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

func UnpadPKCS(a []byte) ([]byte, error) {
	if len(a) == 0 {
		panic("unpadding empty array")
	}
	last := int(a[len(a)-1])
	if last == 0 {
		return nil, fmt.Errorf("bad padding")
	}
	for i := 1; i < last; i++ {
		pos := len(a) - 1 - i
		if int(a[pos]) != last {
			return nil, fmt.Errorf("bad padding")
		}
	}
	return a[:len(a)-last], nil
}
