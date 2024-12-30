package concurrency

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAtomicInsteadLock(t *testing.T) {
	var c atomic.Uint64

	fn := func() {
		t.Log("call")
	}

	const cnt = 10
	var wg sync.WaitGroup
	wg.Add(cnt)
	for i := 0; i < cnt; i++ {
		go func() {
			defer wg.Done()

			if c.Add(1)%3 == 0 {
				fn()
			}
		}()
	}
	wg.Wait()
	require.Equal(t, int32(10), c.Load())
}
