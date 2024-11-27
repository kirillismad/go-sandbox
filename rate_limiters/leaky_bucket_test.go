package rate_limiters

import (
	"context"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestLeakyBucketRateLimiter(t *testing.T) {
	t.Parallel()

	t.Run("default", func(t *testing.T) {
		t.Parallel()

		//r := require.New(t)
		a := assert.New(t)

		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		t.Cleanup(cancel)
		limiter := NewLeakyBucketRateLimiter(
			ctx,
			map[string]CapLimit{},
			WithTTL[*LeakyBucketRateLimiter](defaultTTL),
			WithDefaultLimit[*LeakyBucketRateLimiter](CapLimit{
				Limit: Limit{
					Unit:  time.Second,
					Limit: 10,
				},
				Cap: 8,
			}),
			WithCleanupInterval[*LeakyBucketRateLimiter](defaultCleanupInterval),
		)

		ctx, cancel = context.WithTimeout(ctx, 3*time.Second)
		t.Cleanup(cancel)
		ticker := time.NewTicker(time.Second / 10)
		t.Cleanup(ticker.Stop)

		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					wg.Add(1)
					go func() {
						defer wg.Done()
						_, err := limiter.Acquire(context.Background(), "login", "0.0.0.0")
						a.NoError(err)
					}()
				}
			}
		}()
		wg.Wait()
	})
}
