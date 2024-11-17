package ratelimiters

import (
	"context"
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

type Item struct {
	tokens     int64
	lastAccess time.Time
}

type Limit struct {
	Unit  time.Duration
	Limit int64
}

type TokenBucketRateLimiter struct {
	state        map[string]Item
	m            sync.Mutex
	limits       atomic.Value
	defaultLimit atomic.Value
}

func NewTokenBucketRateLimiter(limits map[string]Limit, defaultLimit Limit) *TokenBucketRateLimiter {
	l := &TokenBucketRateLimiter{
		state: make(map[string]Item),
	}
	l.limits.Store(limits)
	l.defaultLimit.Store(defaultLimit)

	return l
}

func (l *TokenBucketRateLimiter) Acquire(ctx context.Context, operation string, ip string) (Result, error) {
	key := fmt.Sprintf("rate_limit:%s:%s", operation, ip)

	limitMap := l.limits.Load().(map[string]Limit)

	limit, ok := limitMap[operation]
	if !ok {
		limit = l.defaultLimit.Load().(Limit)
	}

	now := time.Now()

	l.m.Lock()
	defer l.m.Unlock()

	item, ok := l.state[key]
	if !ok {
		item = Item{tokens: limit.Limit, lastAccess: now}
	}

	timePassed := now.Sub(item.lastAccess)
	rate := float64(limit.Limit) / float64(limit.Unit.Seconds())

	total := item.tokens + int64(math.Floor(float64(timePassed.Seconds())*rate))
	tokens := min(total, limit.Limit)

	if tokens < 1 {
		return Result{}, NewErrRetryAfter(time.Unix(now.Unix()+int64(math.Floor(1.0/rate)), 0))
	}

	l.state[key] = Item{tokens: tokens - 1, lastAccess: now}

	return Result{
		Remaining: int64(item.tokens),
		Limit:     limit.Limit,
	}, nil
}

func (l *TokenBucketRateLimiter) SetLimits(limits map[string]Limit) {
	l.limits.Store(limits)
}

func (l *TokenBucketRateLimiter) SetDefaultLimit(defaultLimit Limit) {
	l.defaultLimit.Store(defaultLimit)
}
