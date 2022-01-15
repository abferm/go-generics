package gmath_test

import (
	"math"
	"testing"

	"github.com/abferm/go-generics/gmath"
	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	// TODO: add fuzz

	// Positive stays positive
	assert.Equal(t, int(8), gmath.Abs(int(8)))
	assert.Equal(t, float32(8), gmath.Abs(float32(8)))
	assert.Equal(t, math.Inf(1), gmath.Abs(math.Inf(1)))

	// Negative becomes positive
	assert.Equal(t, int(8), gmath.Abs(int(-8)))
	assert.Equal(t, float32(8), gmath.Abs(float32(-8)))
	assert.Equal(t, math.Inf(1), gmath.Abs(math.Inf(-1)))

	// NaN
	assert.Equal(t, float64(1), gmath.Sign(gmath.Abs(math.NaN())))
	assert.Equal(t, float64(1), gmath.Sign(gmath.Abs(gmath.FlipSign(math.NaN()))))
}
