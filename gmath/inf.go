package gmath

import (
	"math"

	"github.com/abferm/go-generics"
)

// Inf returns positive infinity if sign >= 0, negative infinity if sign < 0.
func Inf[F generics.Float, S generics.SignedNumber](sign S) F {
	return F(math.Inf(int(sign)))
}

// IsInf reports whether f is an infinity, according to sign.
// If sign > 0, IsInf reports whether f is positive infinity.
// If sign < 0, IsInf reports whether f is negative infinity.
// If sign == 0, IsInf reports whether f is either infinity.
func IsInf[F generics.Float, S generics.SignedNumber](f F, sign S) bool {
	return math.IsInf(float64(f), int(sign))
}
