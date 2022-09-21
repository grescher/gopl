// Write a function that counts the number of bits that are different in two
// SHA256 hashes. (See PopCount from Section 2.6.2.)
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%d\n", c1, c2, diffBits(c1, c2))
}

// pc[i] is the population count of i.
var pc [sha256.Size * 8]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// diffBits counts the number of bits that are different in two SHA256 digests.
func diffBits(x, y [sha256.Size]byte) (count int) {
	for i := range x {
		count += int(pc[x[i]&^y[i]])
	}
	return count
}
