package challenge4

import (
	"fmt"
	"os"
	"strings"

	c1 "cryptopals/Set1/challenge1"
)

/*
XORed against a single string, find the character
1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736
*/

func charFreq(char byte) float32 {
	wm := map[byte]float32{
		byte('E'): 12.02,
		byte('T'): 9.10,
		byte('A'): 8.12,
		byte('O'): 7.68,
		byte('I'): 7.31,
		byte('N'): 6.95,
		byte('S'): 6.28,
		byte('R'): 6.02,
		byte('H'): 5.92,
		byte('D'): 4.32,
		byte('L'): 3.98,
		byte('U'): 2.88,
		byte('C'): 2.71,
		byte('M'): 2.61,
		byte('F'): 2.30,
		byte('Y'): 2.11,
		byte('W'): 2.09,
		byte('G'): 2.03,
		byte('P'): 1.82,
		byte('B'): 1.49,
		byte('V'): 1.11,
		byte('K'): 0.69,
		byte('X'): 0.17,
		byte('Q'): 0.11,
		byte('J'): 0.10,
		byte('Z'): 0.07,
		byte(' '): 8,
		byte('a'): 8.167,
		byte('b'): 01.492,
		byte('c'): 02.782,
		byte('d'): 04.253,
		byte('e'): 12.702,
		byte('f'): 02.228,
		byte('g'): 02.015,
		byte('h'): 06.094,
		byte('i'): 06.966,
		byte('j'): 00.153,
		byte('k'): 00.772,
		byte('l'): 04.025,
		byte('m'): 02.406,
		byte('n'): 06.749,
		byte('o'): 07.507,
		byte('p'): 01.929,
		byte('q'): 00.095,
		byte('r'): 05.987,
		byte('s'): 06.327,
		byte('t'): 9.056,
		byte('u'): 02.758,
		byte('v'): 00.978,
		byte('w'): 02.360,
		byte('x'): 00.150,
		byte('y'): 01.974,
		byte('z'): 00.074,
	}
	return wm[char]
}

func Challenge4() {
	answer, err := DecodeFromFile("./challenge4/4.txt")
	if err != nil {
		fmt.Printf("Error Encountered: %s", err)
		return
	}
	fmt.Printf("Answer:%s", answer)
}

func DecodeFromFile(file string) ([]byte, error) {
	f, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(f), "\n")
	var answer []byte
	var score float32
	for i, line := range lines {
		ans, tempScore, err := SingleXorCipher([]byte(line))
		if err != nil {
			return nil, err
		}
		fmt.Printf("Checking line %d, Score: %f\n", i, tempScore)
		if tempScore > score {
			answer = ans
			score = tempScore
		}
	}
	return answer, nil
}

func SingleXorCipher(codedMessage []byte) ([]byte, float32, error) {
	raw := c1.DecodeHex(codedMessage)
	var answer []byte
	var score float32
	for i := 0; i < 256; i++ {
		tempAns := make([]byte, len(raw))
		var tempSc float32
		for j := 0; j < len(raw); j++ {
			c := raw[j] ^ byte(i)
			tempSc += charFreq(c)
			tempAns[j] = c
		}
		if tempSc > score {
			answer = tempAns
			score = tempSc
		}
		tempSc = 0
	}
	return answer, score, nil
}
