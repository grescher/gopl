// Package lenconv performs feet and meters conversions.
package lenconv

import "fmt"

type Foot float64
type Meter float64

const (
	MeterPerFoot Meter = 0.3048
	FootPerMeter Foot  = 3.28084
)

func (f Foot) String() string  { return fmt.Sprintf("%.4g ft", f) }
func (m Meter) String() string { return fmt.Sprintf("%.4g m", m) }
