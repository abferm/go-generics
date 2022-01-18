package chantools

import (
	"context"
	"sync"
)

// DynamicFanOut distributes items pushed to the InChan across multiple out channels
// This component requires a separate goroutine to call Run, which will process
// the input channel until either the input channel is closed or the context is
// canceled.
type DynamicFanOut[T any] struct {
	in   <-chan T
	outs map[<-chan T]chan T
	lock sync.Mutex
}

// NewDynamicFanOut produces a DynamicFanOut for the provided input channel
// Close the input channel to signal the end of input
func NewDynamicFanOut[T any](in <-chan T) *DynamicFanOut[T] {
	return &DynamicFanOut[T]{
		in:   in,
		outs: make(map[<-chan T]chan T),
		lock: sync.Mutex{},
	}
}

// Out creates an output channel of the specified size.
// Run will block and wait on unbuffered or full output channels
func (f *DynamicFanOut[T]) Out(size int) <-chan T {
	f.lock.Lock()
	defer f.lock.Unlock()
	out := make(chan T, size)
	f.outs[out] = out
	return out
}

// Close closes the specified output channel and removes it from the Fan
func (f *DynamicFanOut[T]) Close(out <-chan T) {
	// ensure channel keeps draining to prevent deadlock
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
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
	wg.Wait()
}

// Run processes data from the input channel until it is closed, or the provided context is canceled.
// If the input channel is closed the return value will be nil, otherwise it will be ctx.Err().
// All outputs will be closed and removed on completion.
func (f *DynamicFanOut[T]) Run(ctx context.Context) error {
	defer f.closeAll()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case value, ok := <-f.in:
			if !ok {
				// input has closed, we are done
				return nil
			}
			err := f.fan(ctx, value)
			if err != nil {
				return err
			}
		}
	}
}

func (f *DynamicFanOut[T]) fan(ctx context.Context, value T) error {
	f.lock.Lock()
	defer f.lock.Unlock()
	for _, out := range f.outs {
		select {
		case out <- value:
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return nil
}

func (f *DynamicFanOut[T]) closeAll() {
	f.lock.Lock()
	defer f.lock.Unlock()
	for k, v := range f.outs {
		close(v)
		delete(f.outs, k)
	}
}
