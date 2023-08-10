package challenge13

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"log"
	"net/url"

	c7 "cryptopals/Set1/challenge7"
	c11 "cryptopals/Set2/challenge11"
	c9 "cryptopals/Set2/challenge9"
)

var key []byte
var unknown []byte

func init() {
	key = c11.RandBytes(16)
}

func Challenge13() {

	email := "foobar@domain.com"
	ct := profileFor(email)
	fmt.Printf("encoded = %#v\n", hex.EncodeToString(ct))

	// email=foobar%40domain.com+++++++ +++&id=100&role= user
	email = "foobar@domain.com          "
	ct = profileFor(email)
	forge := make([]byte, 0)
	forge = append(forge, ct[:3*16]...)
	end := make([]byte, 0)
	end = append(end, ct[3*16:]...)

	// email=foobar%40d omain.com+++++++ admin&id=100&rol e=user
	email = "foobar@domain.com       admin"
	ct = profileFor(email)
	admin := ct[2*16 : 3*16]

	forge = append(forge, admin...)
	forge = append(forge, end...)

	vals := decodeProfile(forge)
	fmt.Printf("decoded forge = %+v", vals)
}

func profileFor(email string) []byte {
	pars := make(url.Values)
	pars.Add("email", email)
	pars.Add("id", "100")
	pars.Add("role", "user")
	src := []byte(pars.Encode())

	blockSize := len(key)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}
	src = c9.PKCS(src, blockSize)

	dst := make([]byte, len(src))
	c7.EncryptAES_ECB(block, dst, src)
	return dst
}

func decodeProfile(ct []byte) url.Values {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}

	dst, _ := c7.DecryptAES_ECB(ct, []byte(key), block.BlockSize())
	dst, err = c9.UnpadPKCS(dst)
	if err != nil {
		log.Fatal(err)
	}
	vals, err := url.ParseQuery(string(dst))
	if err != nil {
		log.Fatal(err)
	}
	return vals
}
