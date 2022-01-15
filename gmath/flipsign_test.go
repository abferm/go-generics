package gmath_test

import (
	"math"
	"testing"

	"github.com/abferm/go-generics/gmath"
	"github.com/stretchr/testify/assert"
)

func TestFlipsign(t *testing.T) {
	// TODO: add fuzz to happy path cases

	// happy path cases
	oneFloat := gmath.FlipSign(float32(-1))
	assert.Equal(t, float32(1), oneFloat)

	negativeInt := gmath.FlipSign(int16(34))
	assert.Equal(t, int16(-34), negativeInt)

	// Edge cases
	// Infinity
	positiveInf := gmath.FlipSign(math.Inf(-1))
	assert.Equal(t, math.Inf(1), positiveInf)

	negativeInf := gmath.FlipSign(math.Inf(1))
	assert.Equal(t, math.Inf(-1), negativeInf)

	// NaN
	flippedNaN := gmath.FlipSign(math.NaN())
	assert.True(t, math.IsNaN(flippedNaN))
	assert.NotEqual(t, math.Signbit(math.NaN()), math.Signbit(flippedNaN))
}
