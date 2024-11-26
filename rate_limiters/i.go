package rate_limiters

import (
	"context"
	"fmt"
	"time"
)

var (
	ErrRateLimitExceeded = fmt.Errorf("rate limit exceeded")
)

type Result struct {
	Remaining int64
	Limit     int64
}

type ErrRetryAfter struct {
	t time.Time
}

func NewErrRetryAfter(retryAfter time.Time) ErrRetryAfter {
	return ErrRetryAfter{
		t: retryAfter,
	}
}

func (e ErrRetryAfter) Error() string {
	return fmt.Sprintf("retry after %s", e.RetryAfter().Format(time.RFC1123))
}

func (e ErrRetryAfter) RetryAfter() time.Time {
	return e.t
}

type RateLimiter interface {
	Acquire(ctx context.Context, operation, ip string) (Result, error)
}
