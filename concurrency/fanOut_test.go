package concurrency

import (
	"context"
	"slices"
	"sync"
	"testing"
)

func TestFanOut(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	const goCount = 3
	const items = 10
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	outs := FanOut(ctx, PipelineBuffer(ctx, Generator(ctx, items, DigitsGenFunc()), goCount), goCount)

	resultChannel := make(chan int, len(outs))
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(outs))
		for _, out := range outs {
			out := out
			go func() {
				defer wg.Done()
				for e := range out {
					resultChannel <- e
				}
			}()
		}
		wg.Wait()
		close(resultChannel)
	}()

	result := make([]int, 0, items)
	for e := range resultChannel {
		result = append(result, e)
	}
	slices.Sort(result)
	if slices.Compare(want, result) != 0 {
		t.Errorf("got: %v, want: %v", result, want)
	}
}

func TestFanOutCancel1(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch1 := make(chan int)
	outs := FanOut(ctx, ch1, 2)

	cancel()

	for i, out := range outs {
		if _, ok := <-out; ok {
			t.Errorf("channel[%v] is not closed", i)
		}
	}
}

func TestFanOutCancel2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan int, 1)
	ch <- 0
	outs := FanOut(ctx, ch, 2)

	cancel()

	for i, out := range outs {
		if _, ok := <-out; ok {
			t.Errorf("channel[%v] is not closed", i)
		}
	}
}
