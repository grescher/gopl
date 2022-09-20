// Write a function that reports whether two strings are anagrams of each other,
// that is, they contain the same letters in a different order.
package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

// anagrams returns true if two strings are anagrams.
func anagrams(s1, s2 string) bool {
	if utf8.RuneCountInString(s1) != utf8.RuneCountInString(s2) {
		return false
	}
	l1, l2 := countLetters(s1), countLetters(s2)
	for k, v1 := range l1 {
		if v2, ok := l2[k]; !ok {
			return false
		} else if v1 != v2 {
			return false
		}
	}
	return true
}

// countLetters returns the count of runes that appear in the string.
func countLetters(s string) map[rune]int {
	count := make(map[rune]int)
	for _, r := range s {
		count[r]++
	}
	return count
}

func main() {
	if len(os.Args) == 3 {
		fmt.Printf(
			"anagrams(%q, %q) == %t\n",
			os.Args[1], os.Args[2], anagrams(os.Args[1], os.Args[2]),
		)
	} else {
		fmt.Printf("Usage: %v string string\n", os.Args[0])
	}
}
