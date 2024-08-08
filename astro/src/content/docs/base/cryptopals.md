---
title: Cryptopals
description: OverTheWire's Bandit.
---

# [The Cryptopals Crypto Challenges](https://cryptopals.com/)

There are eight sets which get progressively harder. Some of the attacks are clever, but they are not puzzles, they are based off real-world vulnerabilities. The challenges are designed to get you thinking like a cryptographer and to understand how to break cryptography.

-   [Set 1: Basics](https://cryptopals.com/sets/1) (Completed)

-   [Set 2: Block crypto](https://cryptopals.com/sets/2) (Completed)

-   [Set 3: Block and stream crypto](https://cryptopals.com/sets/3)

-   [Set 4: Stream crypto and randomness](https://cryptopals.com/sets/4)

-   [Set 5: Diffie-Hellman and friends](https://cryptopals.com/sets/5)

-   [Set 6: RSA and DSA](https://cryptopals.com/sets/6)

-   [Set 7: Hashes](https://cryptopals.com/sets/7)

-   [Set 8: Abstract Algebra](https://cryptopals.com/sets/8)

I have done the first two sets. I will be doing the rest of the sets in the future. I will be documenting my progress in this repository. Mosts of the set's solution are written in Golang, javascript and Python.

### How to Run

-   Uncomment the specific challenge import and call to the function you want to run in `main.go` at root of project.
-   Run `go run *.go` from root of project. For windwows, run `go run .` from root of project.

### Code

```go
package main

import (

	//c1 "cryptopals/Set1/challenge1"
	//c2 "cryptopals/Set1/challenge2"
	//c3 "cryptopals/Set1/challenge3"
	//c4 "cryptopals/Set1/challenge4"
	//c5 "cryptopals/Set1/challenge5"
	//c6 "cryptopals/Set1/challenge6"
	//c7 "cryptopals/Set1/challenge7"
	//c8 "cryptopals/Set1/challenge8"
	//c9 "cryptopals/Set2/challenge9"
	//c10 "cryptopals/Set2/challenge10"
	//c11 "cryptopals/Set2/challenge11"
	//c12 "cryptopals/Set2/challenge12"
	//c13 "cryptopals/Set2/challenge13"
	//c14 "cryptopals/Set2/challenge14"
	//c15 "cryptopals/Set2/challenge15"
	c16 "cryptopals/Set2/challenge16"
)

func main() {

	// SET 1
	//c1.Challenge1()
	//c2.Challenge2()
	//c3.Challenge3()
	//c4.Challenge4()
	//c5.Challenge5()
	//c6.CryptoPals()
	//c7.Challenge7()
	//c8.Challenge8()

	// SET 2
	//c9.Challenge9()
	//c10.Challenge10()
	//c11.Challenge11()
	//c12.Challenge12()
	//c13.Challenge13()
	//c14.Challenge14()
	//c15.Challenge15()
	c16.Challenge16()
}
```
