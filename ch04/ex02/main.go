// Write a program that prints the SHA256 hash of its standard input by default
// but supports a command-line flag to print the SHA384 or SHA512 hash instead.
package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
	"os"
)

var a = flag.Int("a", 256, "SHA algorithm: 256 (default), 384, 512")

func main() {
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := []byte(scanner.Text())
		switch *a {
		case 256:
			fmt.Printf("SHA256: %x\n", sha256.Sum256(data))
		case 384:
			fmt.Printf("SHA384: %x\n", sha512.Sum384(data))
		case 512:
			fmt.Printf("SHA512: %x\n", sha512.Sum512(data))
		default:
			log.Fatal("digest: unexpected algorithm value")
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("digest: %v\n", err)
	}
}
