package concurrency

import (
	"context"
)

func PipelineBuffer[T any](ctx context.Context, in <-chan T, buf int) <-chan T {
	out := make(chan T, buf)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case e, ok := <-in:
				if !ok {
					return
				}
				select {
				case <-ctx.Done():
					return
				case out <- e:
				}
			}
		}
	}()
	return out
}

func PipelineFunc[T any, R any](ctx context.Context, in <-chan T, f func(T) R) <-chan R {
	out := make(chan R)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case e, ok := <-in:
				if !ok {
					return
				}
				select {
				case <-ctx.Done():
					return
				case out <- f(e):
				}
			}
		}
	}()
	return out
}

func PipelineFilter[T any](ctx context.Context, in <-chan T, f func(T) bool) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case e, ok := <-in:
				if !ok {
					return
				}
				if f(e) {
					select {
					case <-ctx.Done():
						return
					case out <- e:
					}
				}
			}
		}
	}()

	return out
}
