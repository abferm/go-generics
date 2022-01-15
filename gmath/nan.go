package gmath

import (
	"math"

	"github.com/abferm/go-generics"
)

// IsNaN reports whether f is an IEEE 754 ``not-a-number'' value.
func IsNaN[F generics.Float](f F) bool {
	return math.IsNaN(float64(f))
}

// NaN returns an IEEE 754 ``not-a-number'' value.
func NaN[F generics.Float]() F {
	return F(math.NaN())
}
