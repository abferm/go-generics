package gmath

import (
	"math"

	"github.com/abferm/go-generics"
)

// Sign returns a value with the sign
// of x and a magnitude of 1.
func Sign[X generics.SignedNumber](x X) X {
	// We can't do type assertions or switches on generics, so cast to empty interface
	switch interface{}(x).(type) {
	case float32, float64:
		// special case to get the sign of any float, including NaN and Zero
		if math.Signbit(float64(x)) {
			return -1
		}
	default:
		if x < 0 {
			return -1
		}
	}
	return 1
}
