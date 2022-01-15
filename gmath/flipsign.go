package gmath

import "github.com/abferm/go-generics"

// Flipsign returns a value with the magnitude
// of x and the opposite sign.
func FlipSign[X generics.SignedNumber](x X) X {
	return -1 * x
}
