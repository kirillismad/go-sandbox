package rate_limiters

import (
	"context"
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

type Item struct {
	Tokens     int64
	LastAccess time.Time
}

type Limit struct {
	Unit  time.Duration
	Limit int64
}

type TokenBucketRateLimiter struct {
	state           map[string]Item
	m               sync.Mutex
	limits          atomic.Value
	defaultLimit    atomic.Value
	ttl             atomic.Value
	cleanupInterval time.Duration
}

func NewTokenBucketRateLimiter(ctx context.Context, limits map[string]Limit, opts ...Option[*TokenBucketRateLimiter]) *TokenBucketRateLimiter {
	l := &TokenBucketRateLimiter{
		state:           make(map[string]Item),
		cleanupInterval: defaultCleanupInterval,
	}
	l.SetLimits(limits)
	l.SetDefaultLimit(defaultLimit)
	l.SetTTL(defaultTTL)

	for _, o := range opts {
		o(l)
	}

	go l.cleanup(ctx)

	return l
}

func (l *TokenBucketRateLimiter) Acquire(ctx context.Context, operation string, ip string) (Result, error) {
	limitMap := l.limits.Load().(map[string]Limit)

	limit, found := limitMap[operation]
	if !found {
		limit = l.defaultLimit.Load().(Limit)
	}

	now := time.Now()
	key := fmt.Sprintf("rate_limit:%s:%s", operation, ip)

	l.m.Lock()
	defer l.m.Unlock()

	item, found := l.state[key]
	if !found {
		tokens := limit.Limit - 1
		l.state[key] = Item{Tokens: tokens, LastAccess: now}
		return Result{
			Remaining: tokens,
			Limit:     limit.Limit,
		}, nil
	}

	timePassed := now.Sub(item.LastAccess)

	rate := float64(limit.Limit) / float64(limit.Unit)
	replenishedTokens := int64(math.Floor(float64(timePassed) * rate))

	tokens := min(item.Tokens+replenishedTokens, limit.Limit)
	if tokens < 1 {
		return Result{}, NewErrRetryAfter(now.Add(time.Duration(math.Ceil(1.0 / rate))))
	}

	tokens--
	l.state[key] = Item{Tokens: tokens, LastAccess: now}

	return Result{
		Remaining: tokens,
		Limit:     limit.Limit,
	}, nil
}

func (l *TokenBucketRateLimiter) cleanup(ctx context.Context) {
	ticker := time.NewTicker(l.cleanupInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			now := time.Now()

			l.m.Lock()
			for k, v := range l.state {
				if now.Sub(v.LastAccess) > l.ttl.Load().(time.Duration) {
					delete(l.state, k)
				}
			}
			l.m.Unlock()
		}
	}
}

func (l *TokenBucketRateLimiter) SetLimits(limits map[string]Limit) {
	l.limits.Store(limits)
}

func (l *TokenBucketRateLimiter) SetDefaultLimit(defaultLimit Limit) {
	l.defaultLimit.Store(defaultLimit)
}

func (l *TokenBucketRateLimiter) SetTTL(ttl time.Duration) {
	l.ttl.Store(ttl)
}
