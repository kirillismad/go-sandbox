package ratelimiters

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTokenBucketRateLimiter_Acquire(t *testing.T) {
	t.Parallel()

	t.Run("default", func(t *testing.T) {
		t.Parallel()

		r := require.New(t)

		ctx, cancel := context.WithCancel(context.Background())

		t.Cleanup(cancel)
		l := NewTokenBucketRateLimiter(
			ctx,
			map[string]Limit{},
			WithDefaultLimit(Limit{
				Unit:  2 * time.Second,
				Limit: 1,
			}),
		)

		result, err := l.Acquire(context.Background(), "login", "127.0.0.1")
		r.NoError(err)
		r.Equal(Result{
			Remaining: 0,
			Limit:     1,
		}, result)
		time.Sleep(time.Second)

		result, err = l.Acquire(context.Background(), "login", "127.0.0.1")
		r.Error(err)
		r.Equal(Result{}, result)
		time.Sleep(time.Second)

		result, err = l.Acquire(context.Background(), "login", "127.0.0.1")
		r.NoError(err)
		r.Equal(Result{
			Remaining: 0,
			Limit:     1,
		}, result)
		time.Sleep(time.Second)
	})
	t.Run("login", func(t *testing.T) {
		t.Parallel()

		r := require.New(t)

		ctx, cancel := context.WithCancel(context.Background())

		t.Cleanup(cancel)
		l := NewTokenBucketRateLimiter(
			ctx,
			map[string]Limit{
				"login": {
					Unit:  2 * time.Second,
					Limit: 1,
				},
			},
		)

		result, err := l.Acquire(context.Background(), "login", "127.0.0.1")
		r.NoError(err)
		r.Equal(Result{
			Remaining: 0,
			Limit:     1,
		}, result)
		time.Sleep(time.Second)

		result, err = l.Acquire(context.Background(), "login", "127.0.0.1")
		r.Error(err)
		r.Equal(Result{}, result)
		time.Sleep(time.Second)

		result, err = l.Acquire(context.Background(), "login", "127.0.0.1")
		r.NoError(err)
		r.Equal(Result{
			Remaining: 0,
			Limit:     1,
		}, result)
		time.Sleep(time.Second)
	})
}
