package wconv

// KToP converts a weight in kilograms to pounds.
func KToP(k Kilogram) Pound { return Pound(k) * KilogramLb }

// PToK converts a weight in pounds to kilograms.
func PToK(p Pound) Kilogram { return Kilogram(p) * PoundKg }
