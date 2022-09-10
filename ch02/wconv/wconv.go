// Package wconv performs pounds and kilograms conversions.
package wconv

import "fmt"

type Kilogram float64
type Pound float64

const (
	KilogramLb Pound    = 2.20462
	PoundKg    Kilogram = 0.453592
)

func (k Kilogram) String() string { return fmt.Sprintf("%.4g kg", k) }
func (p Pound) String() string    { return fmt.Sprintf("%.4g lb", p) }
