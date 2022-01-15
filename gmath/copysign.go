package gmath

import "github.com/abferm/go-generics"

// Copysign returns a value with the magnitude
// of x and the sign of y.
func Copysign[X, Y generics.SignedNumber](x X, y Y) X {
	signX := Sign(x)
	signY := X(Sign(y))
	// Signs are different, let's flip it
	if signX != signY {
		return FlipSign(x)
	}
	return x
}
