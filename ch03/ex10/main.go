// Write a non-recursive version of comma, using bytes.Buffer instead of string
// concatenation.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer

	n := len(s)
	i := (n-1)%3 + 1
	buf.WriteString(s[:i])
	for ; i < n; i += 3 {
		buf.WriteString("," + s[i:i+3])
	}
	return buf.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(comma(text))
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("comma: %v", err)
	}
}
