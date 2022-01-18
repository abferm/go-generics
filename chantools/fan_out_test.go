package chantools_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/abferm/go-generics/chantools"
	"github.com/stretchr/testify/assert"
)

func TestDynamicFanOut(t *testing.T) {
	in := make(chan int, 5)
	for i := 0; i < 5; i++ {
		in <- i
	}

	fan := chantools.NewDynamicFanOut(in)

	canceled, cancel := context.WithCancel(context.Background())
	cancel()

	err := fan.Run(canceled)
	assert.ErrorIs(t, err, context.Canceled, "Context error should have been returned")

	close(in)
	fan = chantools.NewDynamicFanOut(in)
	err = fan.Run(context.Background())
	assert.NoError(t, err, "Completion should not be an error")

	in = make(chan int, 5)
	for i := 0; i < 5; i++ {
		in <- i
	}

	fan = chantools.NewDynamicFanOut(in)
	runCTX, cancel := context.WithTimeout(context.Background(), time.Second)
	fan.Out(0)

	close(in)

	err = fan.Run(runCTX)
	assert.ErrorIs(t, err, context.DeadlineExceeded, "expected timeout with unserviced output")
	cancel()

	in = make(chan int, 5)
	for i := 0; i < 5; i++ {
		in <- i
	}

	fan = chantools.NewDynamicFanOut(in)
	runCTX, cancel = context.WithTimeout(context.Background(), time.Second)
	out := fan.Out(0)

	close(in)

	fan.Close(out)
	err = fan.Run(runCTX)
	assert.NoError(t, err, "removed output should not block fanout")
	cancel()

	in = make(chan int, 5)
	for i := 0; i < 5; i++ {
		in <- i
	}

	fan = chantools.NewDynamicFanOut(in)
	runCTX, cancel = context.WithTimeout(context.Background(), time.Second)
	out1 := fan.Out(cap(in))
	out2 := fan.Out(cap(in))
	out3 := fan.Out(cap(in))

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := fan.Run(runCTX)
		assert.NoError(t, err, "finished input shouldn't error")
	}()
	for i := 0; i < 5; i++ {
		assert.Equal(t, i, <-out1)
	}

	assert.NotPanics(t, func() { fan.Close(make(<-chan int)) }, "should not panic closing unknown out")

	fan.Close(out2)
	assert.Equal(t, 0, len(out2), "out2 should have drained to prevent deadlocks")
	assert.Equal(t, cap(out3), len(out3), "out3 should be full")
	close(in)
	wg.Wait()
	cancel()

	out3Contents := make([]int, len(out3))
	for i := range out3 {
		out3Contents[i] = i
	}

	assert.Equal(t, []int{0, 1, 2, 3, 4}, out3Contents, "Expected out3 to be closed with entire input as contents")

}
