package challenge6

import (
	"encoding/base64"
	"fmt"
	"os"
	"sort"
	"strings"
)

/*
base64d after being encrypted with repeaating key xor

*/

func Challenge6() {

	// byteSend := []byte("HUIfTQsPAh9PE048GmllH0kcDk4TAQsHThsBFkU2AB4BSWQgVB0dQzNTTmVS")
	// output := DecodeOne(byteSend)
	// fmt.Printf("%v", output)

	lists, err := DecodeBase64File("6.txt")

	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	for i, value := range lists {
		get := findKeySize(value)
		fmt.Printf("%d. Size:%d\n", i, get)
	}

	// a := []byte("this is a test")
	// b := []byte("wokka wokka!!!")
	// ham := hammingDistance(a, b)
	// println(ham)

	// plaintext := []byte(`Hello, World`)
	// value := EncryptXORRepeatingKeyThenBase64(plaintext, []byte("ICE"))
	// fmt.Printf("%s", value)

	// get := findKeySize(output)
	// fmt.Printf("%d", get)
}

func makeByteChunks(kl int, ueb []byte) [][]byte {
	chunks := make([][]byte, kl)
	for i, c := range ueb {
		chunks[i%kl] = append(chunks[i%kl], c)
	}
	return chunks
}

type keyLength struct {
	length   int
	strength float64
}

func findKeySize(ciphertext []byte) int {
	minKeySize := 2
	maxKeySize := 40

	bestKeySize := 0
	bestDistance := float64(len(ciphertext) * 8) // Initialize with a large value

	for keySize := minKeySize; keySize <= maxKeySize; keySize++ {
		if len(ciphertext) < keySize*2 {
			continue // Skip this key size if not enough data for slicing
		}

		firstChunk := ciphertext[:keySize]
		secondChunk := ciphertext[keySize : keySize*2]

		distance := float64(hammingDistance(firstChunk, secondChunk)) / float64(keySize)
		if distance < bestDistance {
			bestDistance = distance
			bestKeySize = keySize
		}
	}

	return bestKeySize
}

func getKeyLengths(input []byte) []keyLength {
	kl := make([]keyLength, 0)
	for len := 2; len < 40; len++ {
		var hd float64
		for i := 0; i < 4; i++ {
			si := len * i
			hd += float64(hammingDistance(input[si:si+len], input[si+len:si+(len*2)]))
		}
		hd = hd / float64(len)
		kl = append(kl, keyLength{strength: hd, length: len})
	}
	sort.Slice(kl, func(i, j int) bool {
		return kl[i].strength < kl[j].strength
	})
	return kl
}

func DecodeBase64File(filename string) (map[int][]byte, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(file), "\n")
	allDecoded := make(map[int][]byte)
	for i, line := range lines {

		byteLine := []byte(line)
		decoded := make([]byte, base64.RawStdEncoding.DecodedLen(len(byteLine)))
		base64.RawStdEncoding.Decode(decoded, byteLine)
		allDecoded[i] = decoded
		//fmt.Printf("\n%d. Line:%v \nDecoded:%v", i, line, decoded)
	}

	return allDecoded, nil
}

func DecodeOne(input []byte) []byte {
	decoded := make([]byte, base64.RawStdEncoding.DecodedLen(len(input)))
	base64.RawStdEncoding.Decode(decoded, input)
	return decoded
}

// The Hamming Distance (or edit distance) is the number of differing bits between two buffers
// This is computed by XORing the two buffers and adding up the bits in the resultant buffer.
func hammingDistance(a, b []byte) int {

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

// func EncryptXORRepeatingKeyThenBase64(payload, key []byte) []byte {
// 	cipher := make([]byte, len(payload))
// 	for i := 0; i < len(payload); i++ {
// 		cipher[i] = payload[i] ^ key[i%len(key)]
// 	}
// 	encoded := make([]byte, base64.RawStdEncoding.EncodedLen((len(cipher))))
// 	base64.RawStdEncoding.Encode(encoded, cipher)
// 	return encoded
// }

// func findKeySizes(ciphertext []byte) []int {
// 	minKeySize := 2
// 	maxKeySize := 40

// 	type keySizeDistance struct {
// 		KeySize  int
// 		Distance float64
// 	}

// 	var probableKeySizes []keySizeDistance

// 	for keySize := minKeySize; keySize <= maxKeySize; keySize++ {
// 		if len(ciphertext) < keySize*2 {
// 			continue // Skip this key size if not enough data for slicing
// 		}

// 		firstChunk := ciphertext[:keySize]
// 		secondChunk := ciphertext[keySize : keySize*2]

// 		distance := float64(hammingDistance(firstChunk, secondChunk)) / float64(keySize)
// 		probableKeySizes = append(probableKeySizes, keySizeDistance{KeySize: keySize, Distance: distance})
// 	}

// 	// Sort the probableKeySizes by distance (ascending order)
// 	sort.Slice(probableKeySizes, func(i, j int) bool {
// 		return probableKeySizes[i].Distance < probableKeySizes[j].Distance
// 	})

// 	// Extract the key sizes from probableKeySizes
// 	var keySizes []int
// 	for _, ks := range probableKeySizes {
// 		keySizes = append(keySizes, ks.KeySize)
// 	}

// 	return keySizes
// }
