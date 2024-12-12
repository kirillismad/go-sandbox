package concurrency

import (
	"context"
	"slices"
	"testing"
	"time"
)

func TestFibonacciGenerator(t *testing.T) {
	cases := []struct {
		name string
		n    int
		want []int
	}{
		{
			name: "1",
			n:    10,
			want: []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ch := FibonacciGenerator(ctx, c.n)
			result := make([]int, 0, c.n)

			for e := range ch {
				result = append(result, e)
			}

			if slices.Compare(result, c.want) != 0 {
				t.Error(result, c.want)
			}
		})
	}
}

func TestEvenGenerator(t *testing.T) {
	cases := []struct {
		name string
		n    int
		want []int
	}{
		{
			name: "1",
			n:    5,
			want: []int{2, 4, 6, 8, 10},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ch := EvenGenerator(ctx, c.n)
			result := make([]int, 0, c.n)

			for e := range ch {
				result = append(result, e)
			}

			if slices.Compare(result, c.want) != 0 {
				t.Error(result, c.want)
			}
		})
	}
}

func TestOddGenerator(t *testing.T) {
	cases := []struct {
		name string
		n    int
		want []int
	}{
		{
			name: "1",
			n:    5,
			want: []int{1, 3, 5, 7, 9},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ch := OddGenerator(ctx, c.n)
			result := make([]int, 0, c.n)

			for e := range ch {
				result = append(result, e)
			}

			if slices.Compare(result, c.want) != 0 {
				t.Error(result, c.want)
			}
		})
	}
}

func TestPrimeGenerator(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g := PrimeGenerator(ctx)
	want := []uint{3, 5, 7, 11, 13, 17, 19, 23, 29, 31}

	const count = 10
	result := make([]uint, count)
	for i := 0; i < count; i++ {
		result[i] = <-g
	}
	cancel()
	if slices.Compare(result, want) != 0 {
		t.Errorf("got: %v, want: %v", result, want)
	}
}

func TestGenerator(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g := Generator(ctx, 3, func() int { return 0 })
	want := []int{0, 0, 0}

	result := make([]int, 0, 3)
	for e := range g {
		result = append(result, e)
	}
	if slices.Compare(result, want) != 0 {
		t.Errorf("want: %v, got: %v", result, want)
	}
}

func TestGeneratorInf(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	g := Generator(ctx, -1, DigitsGenFunc())
	timeout := time.After(200 * time.Millisecond)
	for range g {
		select {
		case <-timeout:
			t.Errorf("channel should be closed")
		default:
		}
	}

}

func TestGeneratorCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g := Generator(ctx, 3, func() int { return 0 })
	cancel()

	time.Sleep(100 * time.Microsecond)
	if _, ok := <-g; ok {
		t.Error("channel g is not closed")
	}
}
