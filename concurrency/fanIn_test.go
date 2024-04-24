package concurrency

import (
	"context"
	"slices"
	"testing"
	"time"
)

func TestFanIn(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	out := PipelineBuffer(ctx, FanIn(ctx, EvenGenerator(ctx, 5), OddGenerator(ctx, 5)), 2)
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result := make([]int, 0, 10)
	for e := range out {
		result = append(result, e)
	}
	slices.Sort(result)

	if slices.Compare(result, want) != 0 {
		t.Errorf("got: %v, want: %v", result, want)
	}
}

func TestFanInCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch1 := make(chan int)
	ch2 := make(chan int)
	out := FanIn(ctx, ch1, ch2)

	cancel()

	if _, ok := <-out; ok {
		t.Error("out channel is not closed")
	}
}

func TestFanInCancel1(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch1 := EvenGenerator(ctx, 2)
	ch2 := OddGenerator(ctx, 2)
	out := FanIn(ctx, ch1, ch2)

	cancel()
	time.Sleep(100 * time.Microsecond)
	if _, ok := <-out; ok {
		t.Error("out channel is not closed")
	}
}
