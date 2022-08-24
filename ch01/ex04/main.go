// Modify `dup2` to print the names of all files in which each duplicated line
// occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		checkDups(os.Stdin)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			checkDups(f)
			f.Close()
		}
	}
}

func checkDups(f *os.File) {
	counts := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			fmt.Printf("%s\n", f.Name())
			break
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
