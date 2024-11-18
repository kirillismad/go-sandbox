package rate_limiters

import (
	"errors"
	"fmt"
	"time"
)

var ErrTryLater = errors.New("try later")

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
	Acquire() (Result, error)
}
