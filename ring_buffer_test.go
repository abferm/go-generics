package generics_test

import (
	"testing"

	"github.com/abferm/go-generics"
	"github.com/stretchr/testify/assert"
)

func TestRingBuffer(t *testing.T) {
	const cap = 5
	_, err := generics.NewRingBuffer[int](0)
	assert.Error(t, err, "expected error for capacity 0")

	buffer, err := generics.NewRingBuffer[int](cap)
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, 0, buffer.Size(), "should be empty")
	assert.Equal(t, cap, buffer.Cap(), "should have specified capacity")

	assert.Len(t, buffer.Slice(), 0)

	buffer.Put(1)
	assert.Equal(t, 1, buffer.Size(), "should no longer be empty")
	assert.Equal(t, cap, buffer.Cap(), "should have specified capacity")

	for i := 0; i <= 21; i++ {
		buffer.Put(i)
	}
	assert.Equal(t, cap, buffer.Size(), "buffer should be full")
	assert.Equal(t, cap, buffer.Cap(), "should have specified capacity")

	assert.Equalf(t, []int{17, 18, 19, 20, 21}, buffer.Slice(), "buffer should contain the last %d elements put into it", cap)
}
