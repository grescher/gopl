// Package tempconv performs Celsius, Fahrenheit and Kelvin conversions.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	FreezingK     Kelvin  = 273.15
)

func (c Celsius) String() string    { return fmt.Sprintf("%.4g °C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.4g °F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%.4g K", k) }
