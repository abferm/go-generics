package gmath

import (
	"math"

	"github.com/abferm/go-generics"
)

// Floor returns the greatest integer value less than or equal to x.
//
// Special cases are:
//	Floor(±0) = ±0
//	Floor(±Inf) = ±Inf
//	Floor(NaN) = NaN
func Floor[F generics.Float](x F) F {
	return F(math.Floor(float64(x)))
}

// Ceil returns the least integer value greater than or equal to x.
//
// Special cases are:
//	Ceil(±0) = ±0
//	Ceil(±Inf) = ±Inf
//	Ceil(NaN) = NaN
func Ceil[F generics.Float](x F) F {
	return F(math.Ceil(float64(x)))
}
