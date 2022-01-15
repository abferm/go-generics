package chantools

import (
	"context"
	"sync"
)

// FanOut distributes items pushed to the InChan across multiple out channels
// This component requires a separate goroutine to call Run, which will process
// the input channel until either the input channel is closed or the context is
// canceled.
type FanOut[T any] struct {
	in   <-chan T
	outs map[<-chan T]chan T
	lock sync.Mutex
}

// NewFanOut produces a FanOut for the provided input channel
// Close the input channel to signal the end of input
func NewFanOut[T any](in <-chan T) *FanOut[T] {
	return &FanOut[T]{
		in:   in,
		outs: make(map[<-chan T]chan T),
		lock: sync.Mutex{},
	}
}

// Out creates an output channel of the specified size.
// Run will block and wait on unbuffered or full output channels
func (f *FanOut[T]) Out(size int) <-chan T {
	f.lock.Lock()
	defer f.lock.Unlock()
	out := make(chan T, size)
	f.outs[out] = out
	return out
}

// Close closes the specified output channel and removes it from the Fan
func (f *FanOut[T]) Close(out <-chan T) {
	// ensure channel keeps draining to prevent deadlock
	go func() {
		for range out {
		}
	}()

	f.lock.Lock()
	defer f.lock.Unlock()
	fullChan, ok := f.outs[out]
	if !ok {
		return
	}
	delete(f.outs, out)
	close(fullChan)
}

// Run processes data from the input channel until it is closed, or the provided context is canceled.
// If the input channel is closed the return value will be nil, otherwise it will be ctx.Err().
// All outputs will be closed and removed on completion.
func (f *FanOut[T]) Run(ctx context.Context) error {
	defer f.closeAll()
	for {
		err := f.itterate(ctx)
		if err != nil {
			return err
		}
	}
}

func (f *FanOut[T]) itterate(ctx context.Context) error {
	f.lock.Lock()
	defer f.lock.Unlock()
	select {
	case v, ok := <-f.in:
		if !ok {
			// input has closed, we are done
			return nil
		}
		for _, out := range f.outs {
			select {
			case out <- v:
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}

func (f *FanOut[T]) closeAll() {
	f.lock.Lock()
	defer f.lock.Unlock()
	for k, v := range f.outs {
		close(v)
		delete(f.outs, k)
	}
}
