//go:build redis

package rate_limiters_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"sandbox/rate_limiters"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

func init() {
	redisHost = os.Getenv("REDIS_HOST")
	if redisHost == "" {
		log.Fatal("REDIS_HOST is not set")
	}
	redisPort = os.Getenv("REDIS_PORT")
	if redisPort == "" {
		log.Fatal("REDIS_PORT is not set")
	}
}

var redisHost string
var redisPort string

func TestRedisTokenBucketRateLimiter_Acquire(t *testing.T) {
	t.Parallel()

	t.Run("default", func(t *testing.T) {
		r := require.New(t)

		ctx := context.Background()

		rdb := redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
		})

		pingResult, err := rdb.Ping(ctx).Result()
		r.NoError(err)
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

		result, err := l.Acquire(ctx, "login", "127.0.0.1")
		r.NoError(err)
		r.Equal(rate_limiters.Result{
			Remaining: 0,
			Limit:     1,
		}, result)
		time.Sleep(time.Second)

		result, err = l.Acquire(ctx, "login", "127.0.0.1")
		r.Error(err)
		r.Equal(rate_limiters.Result{}, result)
		time.Sleep(time.Second)

		result, err = l.Acquire(ctx, "login", "127.0.0.1")
		r.NoError(err)
		r.Equal(rate_limiters.Result{
			Remaining: 0,
			Limit:     1,
		}, result)
		time.Sleep(time.Second)
	})
}
