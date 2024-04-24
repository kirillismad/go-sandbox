package concurrency

import (
	"context"
	"sync"
)

func FanIn[T any](ctx context.Context, ins ...<-chan T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		var wg sync.WaitGroup
		wg.Add(len(ins))

		for i := range ins {
			in := ins[i]
			go func() {
				defer wg.Done()
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
		wg.Wait()
	}()

	return out
}
