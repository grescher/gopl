// Write a general-purpose unit-conversion program analogous to cf that reads
// numbers from its command-line arguments of from the standard input if there
// are no arguments, and converts each number into units like temperature in
// Celsius and Fahrenheit, length in feet and meters, weight in pounds and
// kilograms, and the like.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/grescher/gopl/ch02/lenconv"
	"github.com/grescher/gopl/ch02/tempconv"
	"github.com/grescher/gopl/ch02/wconv"
)

var (
	temp   = flag.Bool("t", false, "Print temperature conversion.")
	length = flag.Bool("l", false, "Print length conversion.")
	weight = flag.Bool("w", false, "Print weight conversion.")
)

func main() {
	flag.Parse()

	// if no flag is set, print all conversion types.
	if flag.NFlag() < 1 {
		*temp, *length, *weight = true, true, true
	}

	// if no command-line arguments were set, read from the standard input.
	args := flag.Args()
	if len(args) < 1 {
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		args = strings.Fields(input.Text())
		if err := input.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
	}

	// print selected conversion type for each argument (cli or stdin).
	for _, arg := range args {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		if *temp {
			printTemp(t)
		}
		if *length {
			printLength(t)
		}
		if *weight {
			printWeight(t)
		}
	}
}

// printTemp prints temperature conversion.
func printTemp(t float64) {
	c := tempconv.Celsius(t)
	f := tempconv.Fahrenheit(t)
	fmt.Printf("%s = %s, %s = %s\n",
		c, tempconv.CToF(c), f, tempconv.FToC(f))
}

// printLength prints length conversion.
func printLength(t float64) {
	f := lenconv.Foot(t)
	m := lenconv.Meter(t)
	fmt.Printf("%s = %s, %s = %s\n",
		f, lenconv.FToM(f), m, lenconv.MToF(m))
}

// printWeight prints weight conversion.
func printWeight(t float64) {
	k := wconv.Kilogram(t)
	lb := wconv.Pound(t)
	fmt.Printf("%s = %s, %s = %s\n",
		k, wconv.KToP(k), lb, wconv.PToK(lb))
}
