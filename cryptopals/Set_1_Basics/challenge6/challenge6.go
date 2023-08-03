package challenge6

import (
	"fmt"
)

func Challenge6() {
	a := []byte("this is a test")
	b := []byte("wokka wokka!!!")

	dist := HammingDistance(a, b)
	fmt.Printf("%d", dist)
}

// The Hamming Distance (or edit distance) is the number of differing bits between two buffers
// This is computed by XORing the two buffers and adding up the bits in the resultant buffer.
func HammingDistance(a, b []byte) int {

	diff := 0
	for i := range a {

		for j := 0; j < 8; j++ {
			mask := byte(1 << uint(j))

			if (a[i] & mask) != (b[i] & mask) {
				diff++
			}
		}
	}
	return diff
}
