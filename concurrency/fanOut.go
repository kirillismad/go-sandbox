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

func FanOutSimple[T any](in <-chan T, outCnt int) []chan T {
	outs := make([]chan T, 0, outCnt)
	for range outCnt {
		outs = append(outs, make(chan T))
	}

	for _, ch := range outs {
		go func() {
			defer close(ch)
			for e := range in {
				ch <- e
			}
		}()
	}

	return outs
}
