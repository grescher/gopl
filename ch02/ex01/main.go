package main

import (
	"fmt"

	"github.com/grescher/gopl/ch02/tempconv"
)

func main() {
	fmt.Printf("Brrrrrr! %v\n", tempconv.KToF(tempconv.Kelvin(0)))
	fmt.Printf("Still brrrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Printf("Ice is melting! %v\n\n", tempconv.FreezingK)
	fmt.Println(tempconv.BoilingC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
	fmt.Println(tempconv.CToK(tempconv.BoilingC))
	fmt.Println("Water is boiling! Let's make some tea")
}
