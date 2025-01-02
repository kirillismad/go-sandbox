package concurrency

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAtomicInsteadLock(t *testing.T) {
	t.Parallel()

	var c atomic.Uint64
	var res atomic.Int64
	fn := func() {
		res.Add(1)
	}

	const cnt = 30
	const div = 3
	var wg sync.WaitGroup
	wg.Add(cnt)
	for i := 0; i < cnt; i++ {
		go func() {
			defer wg.Done()

			if c.Add(1)%div == 0 {
				fn()
			}
		}()
	}
	wg.Wait()
	require.Equal(t, int64(cnt/div), res.Load())
}
