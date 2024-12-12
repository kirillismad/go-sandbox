package concurrency

import (
	"context"
	"fmt"
	"slices"
	"sync"
	"testing"
)

func TestTee(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	in := Generator(ctx, 5, DigitsGenFunc())
	want := []int{0, 1, 2, 3, 4}

	const count = 2
	outs := Tee(ctx, in, count)
	results := make([][]int, count)

	var wg sync.WaitGroup
	for i, ch := range outs {
		i, ch := i, ch
		wg.Add(1)
		go func() {
			defer wg.Done()
			for e := range ch {
				results[i] = append(results[i], e)
			}
		}()
	}
	wg.Wait()

	fmt.Printf("results: %v\n", results)
	for _, result := range results {
		if slices.Compare(result, want) != 0 {
			t.Errorf("got: %v, want: %v", result, want)
		}
	}
}
