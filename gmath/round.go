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

// Trunc returns the integer value of x.
//
// Special cases are:
//	Trunc(±0) = ±0
//	Trunc(±Inf) = ±Inf
//	Trunc(NaN) = NaN
func Trunc[F generics.Float](x F) F {
	return F(math.Trunc(float64(x)))
}

// Round returns the nearest integer, rounding half away from zero.
//
// Special cases are:
//	Round(±0) = ±0
//	Round(±Inf) = ±Inf
//	Round(NaN) = NaN
func Round[F generics.Float](x F) F {
	return F(math.Round(float64(x)))
}
