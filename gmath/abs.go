package gmath

import "github.com/abferm/go-generics"

// Abs returns the absolute value of x.
//
// Special cases are:
//	Abs(±Inf) = +Inf
//	Abs(NaN) = NaN
func Abs[N generics.SignedNumber](n N) N {
	return Copysign(n, 1)
}
