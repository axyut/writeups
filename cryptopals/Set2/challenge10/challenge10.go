package challenge10

import (
	//"bytes"
	"crypto/aes"
	"crypto/cipher"
	c6 "cryptopals/Set1/challenge6"
	"fmt"
	"log"
)

func Challenge10() {

	ct, err := c6.Decodebase64LineByLine("./Set2/challenge10/10.txt")
	if err != nil {
		log.Fatal(err)
	}
	pt := make([]byte, len(ct))
	key := "YELLOW SUBMARINE"
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}
	IV := make([]byte, len(key))
	Decrypt_CBC(block, IV, pt, ct)
	fmt.Printf("pt = %s\n", string(pt))

	// check that if we encrypt and then descrypt the result we would get the
	// same text
	// CBCEncrypt(block, IV, ct, pt)
	// pt2 := make([]byte, len(ct))
	// CBCDecrypt(block, IV, pt2, ct)
	// fmt.Printf("pt2 = %#v\n", string(pt2))

	// fmt.Println(bytes.Compare(pt, pt2) == 0)
}

func Encrypt_CBC(block cipher.Block, IV, dst, src []byte) {
	size := block.BlockSize()
	if size != len(IV) {
		panic("IV length should be equal to block size")
	}
	if len(dst) != len(src) {
		panic("length of dst and src should be equal")
	}
	if len(dst)%size != 0 {
		panic("dst should be padded to the block size")
	}
	if len(src)%size != 0 {
		panic("src should be padded to the block size")
	}
	prev := IV
	for i := 0; i < len(src)/size; i++ {
		x := size * i
		y := size * (i + 1)
		copy(dst[x:y], src[x:y])
		XorBytesInplace(dst[x:y], prev)
		block.Encrypt(dst[x:y], dst[x:y])
		prev = dst[x:y]
	}
}

func Decrypt_CBC(block cipher.Block, IV, dst, src []byte) {
	size := block.BlockSize()
	if size != len(IV) {
		panic("IV length should be equal to block size")
	}
	if len(dst) != len(src) {
		panic("length of dst and src should be equal")
	}
	if len(dst)%size != 0 {
		panic("dst should be padded to the block size")
	}
	if len(src)%size != 0 {
		panic("src should be padded to the block size")
	}
	prev := IV
	for i := 0; i < len(src)/size; i++ {
		x := size * i
		y := size * (i + 1)
		copy(dst[x:y], src[x:y])
		block.Decrypt(dst[x:y], dst[x:y])
		XorBytesInplace(dst[x:y], prev)
		prev = src[x:y]
	}
}

// XorBytesInplace xors bytes from dst and src and puts the result into dst.
func XorBytesInplace(dst, src []byte) {
	tempSrc := src
	if len(dst) < len(src) {
		tempSrc = dst
	}
	for i := range tempSrc {
		dst[i] = dst[i] ^ src[i]
	}
}
