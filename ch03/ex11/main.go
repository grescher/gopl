// Enhance comma so that it deals correctly with floating-point numbers and an
// optional sign.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

// comma inserts commas in a numeric string.
func comma(s string) string {
	var buf bytes.Buffer
	if strings.HasPrefix(s, "-") {
		buf.WriteRune('-')
		s = s[1:]
	}

	var n, dot int
	if dot = strings.LastIndex(s, "."); dot >= 0 {
		n = len(s[:dot])
	} else {
		n = len(s)
	}

	i := (n-1)%3 + 1
	buf.WriteString(s[:i])
	for ; i < n; i += 3 {
		buf.WriteString("," + s[i:i+3])
	}

	if dot >= 0 {
		buf.WriteString(s[dot:])
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
