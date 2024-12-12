package concurrency

import (
	"context"
	"math"
)

func FibonacciGenerator(ctx context.Context, n int) <-chan int {
	return Generator(ctx, n, FibonacciGenFunc())
}

func EvenGenerator(ctx context.Context, n int) <-chan int {
	return Generator(ctx, n, EvenGenFunc())
}

func OddGenerator(ctx context.Context, n int) <-chan int {
	return Generator(ctx, n, OddGenFunc())
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	// Check divisibility for numbers of the form 6k ± 1 up to the square root of n
	for i := 5; i <= int(math.Sqrt(float64(n))); i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

func PrimeGenerator(ctx context.Context) <-chan uint {
	odds := OddGenerator(ctx, -1)
	primes := PipelineFilter(ctx, odds, func(i int) bool { return isPrime(i) })
	out := PipelineFunc(ctx, primes, func(x int) uint { return uint(x) })
	return out
}

type GenFunc[T any] func() T

func DigitsGenFunc() GenFunc[int] {
	i := 0
	return func() int {
		defer func() { i++ }()
		return i
	}
}

func EvenGenFunc() GenFunc[int] {
	i := 1
	return func() int {
		defer func() { i++ }()
		return i * 2
	}
}

func OddGenFunc() GenFunc[int] {
	i := 0
	return func() int {
		defer func() { i++ }()
		return i*2 + 1
	}
}

func FibonacciGenFunc() GenFunc[int] {
	f1, f2 := 0, 1
	var next int
	return func() int {
		next, f1, f2 = f1, f2, f1+f2
		return next
	}
}

func Generator[T any](ctx context.Context, n int, f GenFunc[T]) <-chan T {
	var cnt float64
	if n == -1 {
		cnt = math.Inf(1)
	} else {
		cnt = float64(n)
	}
	result := make(chan T)
	go func() {
		defer close(result)
		for i := 0; float64(i) < cnt; i++ {
			select {
			case <-ctx.Done():
				return
			case result <- f():
			}
		}
	}()
	return result
}
