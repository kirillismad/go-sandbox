package ratelimiters

import (
	"context"
	"testing"
	"time"
)

func TestTokenBucketRateLimiter_Acquire(t *testing.T) {
	t.Parallel()

	l := NewTokenBucketRateLimiter(map[string]Limit{}, Limit{
		Unit:  3 * time.Second,
		Limit: 1,
	})

	for i := 0; i < 10; i++ {
		r, err := l.Acquire(context.Background(), "login", "127.0.0.1")
		t.Logf("result: %#v, err: %v", r, err)
		time.Sleep(time.Second)
	}
}
