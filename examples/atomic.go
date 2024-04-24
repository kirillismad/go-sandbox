package examples

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func DemoAtomic() {
	demoAtomic()
}

func demoAtomic() {
	const count = 1000
	var cnt atomic.Int64
	cnt.Store(0)
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cnt.Add(1)
		}()
	}
	wg.Wait()
	fmt.Printf("cnt: %v\n", cnt.Load())
}
