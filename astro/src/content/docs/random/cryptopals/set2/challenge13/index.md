---
title: "Challenge 13: ECB cut-and-paste"
description: Set 2 of Cryptopals
---

### Probelm

```plaintext
ECB cut-and-paste
Write a k=v parsing routine, as if for a structured cookie. The routine should take:

foo=bar&baz=qux&zap=zazzle
... and produce:

{
  foo: 'bar',
  baz: 'qux',
  zap: 'zazzle'
}
(you know, the object; I don't care if you convert it to JSON).

Now write a function that encodes a user profile in that format, given an email address. You should have something like:

profile_for("foo@bar.com")
... and it should produce:

{
  email: 'foo@bar.com',
  uid: 10,
  role: 'user'
}
... encoded as:

email=foo@bar.com&uid=10&role=user
Your "profile_for" function should not allow encoding metacharacters (& and =). Eat them, quote them, whatever you want to do, but don't let people set their email address to "foo@bar.com&role=admin".

Now, two more easy functions. Generate a random AES key, then:

Encrypt the encoded user profile under the key; "provide" that to the "attacker".
Decrypt the encoded user profile and parse it.
Using only the user input to profile_for() (as an oracle to generate "valid" ciphertexts) and the ciphertexts themselves, make a role=admin profile.
```

### Code

```go
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
```
