package concurrency

import "context"

func Tee[T any](ctx context.Context, in <-chan T, outCnt int) []<-chan T {
	channels := make([]chan T, 0, outCnt)
	for i := 0; i < outCnt; i++ {
		channels = append(channels, make(chan T))
	}
	go func() {
		defer func() {
			for _, ch := range channels {
				close(ch)
			}
		}()
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				for i := range channels {
					ch := channels[i]
					select {
					case <-ctx.Done():
						return
					case ch <- v:
					}
				}
			}
		}
	}()

	outp := make([]<-chan T, len(channels))
	for i, ch := range channels {
		outp[i] = ch
	}
	return outp
}
