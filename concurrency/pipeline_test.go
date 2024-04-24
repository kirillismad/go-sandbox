package concurrency

import (
	"context"
	"fmt"
	"math"
	"sandbox/utils"
	"slices"
	"testing"
	"time"
)

func pow2(n int) float64 {
	return math.Pow(float64(2), float64(n))
}

func TestPipelineFuncOk(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	in := make(chan int)
	pipe := PipelineFunc(ctx, in, pow2)
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	go func() {
		defer close(in)
		for _, v := range s {
			in <- v
		}
	}()

	result := make([]float64, 0, len(s))
	for e := range pipe {
		result = append(result, e)
	}
	if slices.Compare(result, utils.SliceMap(s, pow2)) != 0 {
		t.Errorf("s: %v, result: %v", s, result)
	}
}

func TestPipelineFuncCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pipe := PipelineFunc(ctx, make(chan int), pow2)
	cancel()

	select {
	case <-time.After(100 * time.Millisecond):
		t.Error("There is Lock")
	case _, ok := <-pipe:
		if ok {
			t.Error("Error")
		}
	}
}

func TestPipelineFuncCancel2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	in := make(chan int)
	pipe := PipelineFunc(ctx, in, pow2)
	in <- 42

	cancel()

	select {
	case <-time.After(100 * time.Millisecond):
		t.Error("There is Lock")
	case _, ok := <-pipe:
		if ok {
			t.Error("Error")
		}
	}
}

func TestPipelineOk(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := PipelineBuffer(ctx, FibonacciGenerator(ctx, 10), 3)
	for e := range ch {
		fmt.Printf("e: %v, len: %v\n", e, len(ch))
	}
}

func TestPipelineCancelContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := PipelineBuffer(ctx, make(chan int), 3)

	cancel()

	for e := range ch {
		fmt.Printf("e: %v, len: %v\n", e, len(ch))
	}
}

func TestPipelineCancelContext2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := PipelineBuffer(ctx, FibonacciGenerator(ctx, 4), 3)
	for len(ch) != 3 {
		time.Sleep(100 * time.Millisecond)
	}
	cancel()
}
