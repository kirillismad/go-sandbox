package concurrency

import (
	"context"
)

func FanOut[T any](ctx context.Context, in <-chan T, outCnt int) []chan T {
	outs := make([]chan T, outCnt)
	for i := range outs {
		outs[i] = make(chan T)
	}

	for i := range outs {
		out := outs[i]
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
	}

	return outs
}

func FanOutBroadcast[T any](ctx context.Context, in <-chan T, outCnt int) []chan T {
	outs := make([]chan T, outCnt)
	for i := range outs {
		outs[i] = make(chan T)
	}

	go func() {
		defer func() {
			for _, out := range outs {
				close(out)
			}
		}()
		for {
			select {
			case <-ctx.Done():
				return
			case e, ok := <-in:
				if !ok {
					return
				}
				for _, out := range outs {
					select {
					case <-ctx.Done():
						return
					case out <- e:
					}
				}
			}
		}
	}()

	return outs
}
