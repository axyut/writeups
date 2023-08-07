package challenge6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"

	c1 "cryptopals/Set1/challenge1"
)

type scoredKeysize struct {
	dist    float64
	keySize int
}

func Challenge6() {
	a := []byte("this is a test")
	b := []byte("wokka wokka!!!")

	dist := HammingDistance(a, b)
	fmt.Printf("%d", dist)

	fileBytes, err := Decodebase64LineByLine("./challenge6/6.txt")
	if err != nil {
		log.Fatal(err)
	}

	scoredKeysizes := GetBestKeySize(fileBytes)

	// challenge says 2,3 smallest keysizes, upto 5 should be enough
	for i := range scoredKeysizes[:5] {
		keysize := scoredKeysizes[i].keySize
		fmt.Printf("\nFor %d. Keysize: %d\n", i+1, keysize)

		blocks := TransposeBytes(fileBytes, keysize)

		for i, eachKeySizeBlock := range blocks {
			// TODO send each key size block and figure out the key
			fmt.Printf("%d, %v", i, eachKeySizeBlock)
		}
	}
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

func Decodebase64LineByLine(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := make([]byte, len(scanner.Bytes()))
	for scanner.Scan() {

		byteSize := []byte(scanner.Text())
		decodedLine := c1.DecodeBase64(byteSize)

		result = append(result, decodedLine...)
	}
	if err := scanner.Err(); err != nil {
		return result, err
	}

	return result, nil
}

func GetBestKeySize(file []byte) []scoredKeysize {
	const blocks = 4
	scoredKeysizes := make([]scoredKeysize, 0, 40)

	for keySize := 1; keySize <= 40; keySize++ {
		dist := 0.0

		for i := 0; i < keySize*blocks; i += keySize {
			startEndByte := i + keySize
			endEndByte := i + keySize*2
			dist += float64(HammingDistance(file[i:startEndByte], file[startEndByte:endEndByte])) / float64(keySize)
		}

		dist = dist / float64(blocks)
		scoredKeysizes = append(scoredKeysizes, scoredKeysize{dist, keySize})
	}

	// Sort and print
	sort.Slice(scoredKeysizes, func(i, j int) bool {
		return scoredKeysizes[i].dist < scoredKeysizes[j].dist
	})

	fmt.Printf("Best fitting keysizes and their average distance on %d blocks:\n", blocks)
	for i, v := range scoredKeysizes[:5] {
		fmt.Printf("%d. KeySize: %d, Edit Distance: %f\n", i+1, v.keySize, v.dist)
	}
	return scoredKeysizes
}

func TransposeBytes(fileBytes []byte, keySize int) [][]byte {
	byteArray := make([][]byte, keySize)

	for i, singleByte := range fileBytes {
		tillKeySize := i % keySize                                          // generates from 0 to keysize
		byteArray[tillKeySize] = append(byteArray[tillKeySize], singleByte) //[[ no. of keysize bytes ] []]
	}
	return byteArray
}
