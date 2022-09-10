package lenconv

// FToM converts a length in feet to meters.
func FToM(f Foot) Meter { return Meter(f) * MeterPerFoot }

// MToF converts a length in meters to feet.
func MToF(m Meter) Foot { return Foot(m) * FootPerMeter }
