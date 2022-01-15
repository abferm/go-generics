package generics

import "fmt"

type RingBuffer[T any] struct {
	s []T
	i int
}

func NewRingBuffer[T any](cap int) (*RingBuffer[T], error) {
	if cap < 1 {
		return nil, fmt.Errorf("cap must be greater than or equal to 1")
	}
	return &RingBuffer[T]{
		s: make([]T, 0, cap),
		i: 0,
	}, nil
}

func (b *RingBuffer[T]) Put(value T) {
	if b.Size() < b.Cap() {
		b.s = append(b.s, value)
	} else {
		b.s[b.i] = value
	}
	b.i++
	if b.i == cap(b.s) {
		b.i = 0
	}
}

func (b *RingBuffer[T]) PutMultiple(values ...T) {
	for _, v := range values {
		b.Put(v)
	}
}

func (b RingBuffer[T]) Slice() []T {
	return append(b.s[b.i:], b.s[:b.i]...)
}

func (b RingBuffer[T]) Cap() int {
	return cap(b.s)
}

func (b RingBuffer[T]) Size() int {
	return len(b.s)
}
