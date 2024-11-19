//go:build redis

package rate_limiters_test

import (
	"context"
	"sandbox/rate_limiters"
	"sync"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

func TestRedisTokenBucketRateLimiter_Acquire(t *testing.T) {
	t.Parallel()

	t.Run("default", func(t *testing.T) {
		r := require.New(t)

		ctx := context.Background()

		rdb := redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		})

		pingResult, err := rdb.Ping(ctx).Result()
		if err != nil {
			t.Skip()
		}
		t.Log(pingResult)

		err = rdb.FlushDB(ctx).Err()
		r.NoError(err)
		t.Log("FlushDB")
		err = rdb.ScriptFlush(ctx).Err()
		r.NoError(err)
		t.Log("ScriptFlush")

		l := rate_limiters.NewRedisTokenBucketRateLimiter(
			rdb,
			make(map[string]rate_limiters.Limit),
			rate_limiters.WithDefaultLimit[*rate_limiters.RedisTokenBucketRateLimiter](rate_limiters.Limit{
				Unit:  2 * time.Second,
				Limit: 1,
			}),
		)
		err = l.Prepare(ctx)
		r.NoError(err)

		result, err := l.Acquire(context.Background(), "login", "127.0.0.1")
		r.NoError(err)
		r.Equal(rate_limiters.Result{
			Remaining: 0,
			Limit:     1,
		}, result)
		time.Sleep(time.Second)

		result, err = l.Acquire(context.Background(), "login", "127.0.0.1")
		r.Error(err)
		r.Equal(rate_limiters.Result{}, result)
		time.Sleep(time.Second)

		result, err = l.Acquire(context.Background(), "login", "127.0.0.1")
		r.NoError(err)
		r.Equal(rate_limiters.Result{
			Remaining: 0,
			Limit:     1,
		}, result)
		time.Sleep(time.Second)
	})
	t.Run("parallel", func(t *testing.T) {
		r := require.New(t)

		ctx := context.Background()

		rdb := redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		})

		pingResult, err := rdb.Ping(ctx).Result()
		if err != nil {
			t.Skip()
		}
		t.Log(pingResult)

		err = rdb.FlushDB(ctx).Err()
		r.NoError(err)
		t.Log("FlushDB")

		l := rate_limiters.NewRedisTokenBucketRateLimiter(
			rdb,
			make(map[string]rate_limiters.Limit),
			rate_limiters.WithDefaultLimit[*rate_limiters.RedisTokenBucketRateLimiter](rate_limiters.Limit{
				Unit:  1 * time.Second,
				Limit: 111,
			}),
		)

		var wg sync.WaitGroup

		for i := 0; i < 15; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result, err := l.Acquire(context.Background(), "login", "127.0.0.1")
				r.NoError(err)
				t.Log(result)
			}()
			if i%2 == 1 {
				wg.Add(1)
				go func() {
					defer wg.Done()
					err = rdb.ScriptFlush(ctx).Err()
					r.NoError(err)
					t.Log("ScriptFlush")
				}()
			}
		}
		wg.Wait()
		t.Log(l.Version())
	})
}
