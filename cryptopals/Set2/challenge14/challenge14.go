package challenge14

import (
	"bytes"
	"crypto/aes"
	c6 "cryptopals/Set1/challenge6"
	c7 "cryptopals/Set1/challenge7"
	c11 "cryptopals/Set2/challenge11"
	c9 "cryptopals/Set2/challenge9"
	"fmt"
	"log"
	"os"
)

var key []byte
var prefix []byte
var unknown []byte

func init() {
	key = c11.RandBytes(16)
	prefix = c11.RandBytes(int(c11.RandByte()) % 64)
	var err error
	unknown, err = c6.Decodebase64LineByLine("./Set2/challenge12/12.txt")
	if err != nil {
		panic("Cannot decode the file.")
	}
}

func Challenge14() {
	size := blockSize()
	//fmt.Println("Block size: ", size)

	cipher := EncryptOracle(make([]byte, size*4))
	if !c11.IsECB(cipher, size) {
		panic("Not ECB")
	}
	fmt.Println("ECB detected.")

	prBlocks, prBytes := detectPrefix(size)
	prLen := prBlocks*size + prBytes
	//fmt.Println("Prefix blocks: ", prBlocks, "\nPrefix bytes: ", prBytes)

	cipher = EncryptOracle(make([]byte, 0))
	secrectLen := len(cipher) - prLen
	fmt.Println("Length of the message and padding: ", secrectLen)

	// End prefix with aligned block. We will prepend all inputs with mockup to
	// discard first blocks in the cipher as prefix part.
	mockup := make([]byte, size-prBytes)
	// discard this amount of bytes in cipher
	discard := prLen + len(mockup)

	fill := make([]byte, size)
	unknown := make([]byte, secrectLen)
	for blockLen := 0; blockLen < secrectLen/size; blockLen++ {

		for i := 1; i < 17 && blockLen*size+i-1 < secrectLen; i++ {
			pos := blockLen*size + i - 1

			if pos >= secrectLen {
				break
			}
			if blockLen == 0 {
				copy(fill[len(fill)-i:], unknown[:i])
			} else {
				copy(fill, unknown[pos-size+1:pos])
			}

			cipher := EncryptOracle(append(mockup, fill[:size-i]...))
			ideal := cipher[discard+blockLen*size : discard+(blockLen+1)*size]

			// Determine the last byte by trying all possible values and comparing
			// resulting cipher with ideal.
			found := false
			for b := 0; b < 256; b++ {
				fill[size-1] = byte(b)
				cipher := EncryptOracle(append(mockup, fill...))

				if bytes.Equal(ideal, cipher[discard:discard+size]) {
					unknown[pos] = byte(b)
					found = true
					break
				}
			}
			if !found {
				fmt.Println("Cannot find byte at Position ", pos)
				if len(unknown)-pos < 16 {
					fmt.Printf("Only %d bytes remaining, might be padding.\n", len(unknown)-pos)
				}
				fmt.Printf("Decoded with padding:\n%#v\n", string(unknown))
				os.Exit(1)
			}
		}
	}
	fmt.Printf("Decoded without padding:\n%#v\n", string(unknown))
}

func EncryptOracle(input []byte) []byte {
	src := append(prefix, input...)
	src = append(src, unknown...)

	keySize := len(key)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}
	src = c9.PKCS(src, keySize)
	dst := make([]byte, len(src))
	c7.EncryptAES_ECB(block, dst, src)
	return dst
}

func blockSize() int {
	size := 0
	prevLen := len(EncryptOracle(make([]byte, 1)))

	for i := 2; i < 32; i++ {
		cipher := EncryptOracle(make([]byte, i))

		if prevLen != len(cipher) {
			size = len(cipher) - prevLen
			break
		}
	}
	return size
}

func detectPrefix(size int) (int, int) {
	blocks := 0

	prev := EncryptOracle(make([]byte, 0))
	cipher := EncryptOracle(make([]byte, 1))

	for j := 0; j < len(cipher)-size; j += size {
		if bytes.Equal(prev[j:j+size], cipher[j:j+size]) {
			blocks++
		} else {
			break
		}
	}
	// Prefix contains at least blocks amount of data,
	// now, detect how many more additional bytes. The block after blocks
	// will stop changing when enough bytes were added to the input.
	added := 0
	check := blocks * size

	for i := 1; i <= size+1; i++ {
		input := make([]byte, i)
		cipher := EncryptOracle(input)

		if bytes.Equal(prev[check:check+size], cipher[check:check+size]) {
			break
		}
		added++
		prev = cipher
	}
	return blocks, size - added
}
