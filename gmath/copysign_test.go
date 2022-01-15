package gmath_test

import (
	"math"
	"testing"

	gmath "github.com/abferm/go-generics/gmath"
	"github.com/stretchr/testify/assert"
)

func TestCopySign(t *testing.T) {
	// float with sign of int
	floatWithSignOfInt := gmath.Copysign(float64(42.42), int(-1))
	assert.Equal(t, float64(-42.42), floatWithSignOfInt)
	floatWithSignOfInt = gmath.Copysign(float64(-42.42), int(-1))
	assert.Equal(t, float64(-42.42), floatWithSignOfInt)

	// integer type with sign of positive NaN
	intWithSignOfNaN := gmath.Copysign(int(2), math.NaN())
	assert.Equal(t, int(2), intWithSignOfNaN)
	intWithSignOfNaN = gmath.Copysign(int(-2), math.NaN())
	assert.Equal(t, int(2), intWithSignOfNaN)

	// integer type with sign of negative NaN
	intWithSignOfNegativeNaN := gmath.Copysign(int(16), gmath.FlipSign(math.NaN()))
	assert.Equal(t, int(-16), intWithSignOfNegativeNaN)
	intWithSignOfNegativeNaN = gmath.Copysign(int(-16), gmath.FlipSign(math.NaN()))
	assert.Equal(t, int(-16), intWithSignOfNegativeNaN)

}
