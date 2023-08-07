package challenge6

import (
	"bytes"
	"fmt"
	"log"

	//"encoding/base64"
	"math"
	"math/rand"
	"sort"
	//"strings"
	//c3 "github.com/axyut/cryptopals/challenge3"
)

// Splits a byte slice into length chunks

// Returns a non-cryptographically secure random number between
// `min` and `max` (inclusive)
func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

// Generates `n` random bytes
func GenerateRandomBytes(n int) ([]byte, error) {
	result := make([]byte, n)
	_, err := rand.Read(result)
	if err != nil {
		return []byte{}, err
	}
	return result, nil
}

// Finds a matching block of size `size` in `data`
// Returns the index of the first matching block, -1 if not found
func FindMatchingBlock(data []byte, size int) int {
	chunks := SplitBytes(data, size)

	for i, chunkA := range chunks {
		for j, chunkB := range chunks {
			if i != j && bytes.Equal(chunkA, chunkB) {
				return i * size
			}
		}
	}

	return -1
}
func SplitBytes(buf []byte, length int) [][]byte {
	chunks := [][]byte{}
	for offset := 0; offset < len(buf); offset += length {
		if offset+length >= len(buf) {
			chunks = append(chunks, buf[offset:])
		} else {
			chunks = append(chunks, buf[offset:offset+length])
		}
	}
	return chunks
}

// Transpose a slice of byte slices. e.g. AA,BB,CC == ABC,ABC
func TransposeByte(chunks [][]byte) [][]byte {
	result := make([][]byte, len(chunks[0]))
	for _, chunk := range chunks {
		for j, b := range chunk {
			result[j] = append(result[j], b)
		}
	}
	return result
}

// ///////////// working

func CryptoPals() {
	fileBytes, err := Decodebase64LineByLine("./challenge6/6.txt")
	//fmt.Printf("%d", len(fileBytes))
	if err != nil {
		log.Fatal(err)
	}

	// Find the most likely key sizes. Picking the best one won't do us any
	// good here, we need to check the whole 'hood.
	//const blocks = 4
	//const maxKeysize = 40
	scoredKeysizes := GetBestKeySize(fileBytes)

	// Let's iterate through the potential keysizes
	// challenge says 2,3 smallest keysizes, 5 should be enough
	for i := range scoredKeysizes[:5] {
		keysize := scoredKeysizes[i].keySize
		fmt.Printf("\nFor %d. Keysize: %d\n", i+1, keysize)

		// ct is transposed into keysize blocks. Each chunk is ct xored by one
		// byte. For each chunk we find one byte key.
		blocks := TransposeBytes(fileBytes, keysize)
		//fmt.Printf("%v", blocks)
		guess := make([]byte, keysize)
		avgTopScore := 0.0
		for i, eachKeySizeBlock := range blocks {
			ptScored := BestXorByteKey(eachKeySizeBlock)
			// for _, p := range ptScored[:5] {
			// 	fmt.Printf("%v %#v\n", p.Score, string(p.Text[:40]))
			// }
			guess[i] = ptScored[0].Key
			avgTopScore += ptScored[0].Score
		}
		avgTopScore /= float64(len(blocks))

		fmt.Printf("Avg.Score: %f, Key: %#v\n", avgTopScore, string(guess))
	}
}

func BestXorByteKey(keySizeByte []byte) []ScoredText {
	texts := make([]ScoredText, 0, 256)
	for b := 0; b <= 255; b++ {
		plain := XorBytes(keySizeByte, []byte{byte(b)})
		score := TextScoreEnglish(plain)
		//fmt.Printf("%s", keySizeByte)

		texts = append(texts, ScoredText{score, byte(b), plain})
	}
	sort.Slice(texts, func(i, j int) bool {
		return texts[i].Score > texts[j].Score
	})
	return texts
}
func XorBytes(a, b []byte) []byte {
	if len(a) < len(b) {
		t := a
		a = b
		b = t
	}
	// a is longer than b
	c := make([]byte, len(a))
	for i := range a {
		c[i] = a[i] ^ b[i%len(b)]
	}
	return c
}

type ScoredText struct {
	Score float64
	Key   byte
	Text  []byte
}

func TextScoreEnglish(a []byte) float64 {
	a = bytes.ToLower(a)
	counts := make(map[byte]int)
	for _, ch := range a {
		// if ch < 'a' || ch > 'z' {
		// 	continue
		// }
		if _, ok := counts[ch]; !ok {
			counts[ch] = 0
		}
		counts[ch] += 1
	}
	freqs := make(map[byte]float64)
	for ch, c := range counts {
		freqs[ch] = float64(c) / float64(len(a))
		// fmt.Printf("%c %f\n", ch, freqs[ch])
	}
	return cosineDistance(freqs, FreqLetters)
}

func cosineDistance(x, y map[byte]float64) float64 {
	// x0*y0 + x1*y1 + ..
	// ---------------------------------------------
	// sqroot(x0^2+x1^2+...) + sqroot(y0^1+y1^2+...)

	var upper float64
	var xlen float64
	for xk, xv := range x {
		if yv, ok := y[xk]; ok {
			upper += xv * yv
		}
		xlen += xv * xv
	}
	xlen = math.Sqrt(xlen)

	var ylen float64
	for _, yv := range y {
		ylen += yv * yv
	}
	ylen = math.Sqrt(ylen)

	// fmt.Printf("%f\n-----------------\n%f * %f\n", upper, xlen, ylen)
	return upper / (xlen * ylen)
}

var FreqLetters = map[byte]float64{
	' ':  0.13575645,
	'e':  0.12575645,
	't':  0.09085226,
	'a':  0.08000395,
	'o':  0.07591270,
	'i':  0.06920007,
	'n':  0.06903785,
	's':  0.06340880,
	'h':  0.06236609,
	'r':  0.05959034,
	'd':  0.04317924,
	'l':  0.04057231,
	'u':  0.02841783,
	'c':  0.02575785,
	'm':  0.02560994,
	'f':  0.02350463,
	'w':  0.02224893,
	'g':  0.01982677,
	'y':  0.01900888,
	'p':  0.01795742,
	'b':  0.01535701,
	'v':  0.00981717,
	'k':  0.00739906,
	'x':  0.00179556,
	'j':  0.00145188,
	'q':  0.00117571,
	'z':  0.00079130,
	',':  0.00099130,
	';':  0.00099130,
	'\'': 0.00079130,
	'.':  0.00079130,
}

// func ReadBase64File(filename string) ([]byte, error) {
// 	raw, err := os.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	scanner := bufio.NewScanner(bytes.NewReader(raw))

// 	var sb strings.Builder
// 	for scanner.Scan() {
// 		sb.WriteString(scanner.Text())
// 	}
// 	data := sb.String()

// 	dst := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
// 	n, err := base64.StdEncoding.Decode(dst, []byte(data))
// 	if err != nil {
// 		fmt.Println("decode error:", err)
// 		return nil, err
// 	}
// 	dst = dst[:n]
// 	return dst, nil
// }
