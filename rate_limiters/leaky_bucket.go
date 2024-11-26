package rate_limiters

import (
	"context"
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

// TODO
// 1. Add tests
// 2. Add `cap` option instead creating a bucket with `.Limit` capacity

type request struct {
	done chan struct{}
}

type lbItem struct {
	bucket     chan request
	cancel     context.CancelFunc
	lastAccess time.Time
}

type LeakyBucketRateLimiter struct {
	state           map[string]lbItem
	m               sync.Mutex
	limits          atomic.Value
	defaultLimit    atomic.Value
	ttl             atomic.Value
	cleanupInterval time.Duration
}

func NewLeakyBucketRateLimiter(ctx context.Context, limits map[string]Limit, opts ...Option[*LeakyBucketRateLimiter]) *LeakyBucketRateLimiter {
	l := &LeakyBucketRateLimiter{
		state:           make(map[string]lbItem),
		cleanupInterval: defaultCleanupInterval,
	}

	l.SetTTL(defaultTTL)
	l.SetDefaultLimit(defaultLimit)
	l.SetLimits(limits)

	for _, opt := range opts {
		opt(l)
	}

	go l.cleanup(ctx)

	return l
}

func (l *LeakyBucketRateLimiter) Acquire(ctx context.Context, operation, ip string) (Result, error) {
	limitMap := l.limits.Load().(map[string]Limit)

	limit, found := limitMap[operation]
	if !found {
		limit = l.defaultLimit.Load().(Limit)
	}

	key := fmt.Sprintf("rate_limit:%s:%s", operation, ip)

	l.m.Lock()
	item, found := l.state[key]
	if !found {
		newCtx, cancel := context.WithCancel(context.Background())
		item = lbItem{bucket: make(chan request, limit.Limit), cancel: cancel}

		startLeaker(newCtx, item.bucket, limit)
		l.state[key] = item
	}
	l.m.Unlock()

	if len(item.bucket) >= int(limit.Limit) {
		return Result{}, ErrRateLimitExceeded
	}

	req := request{done: make(chan struct{})}

	select {
	case item.bucket <- req:
		item.lastAccess = time.Now()
	case <-ctx.Done():
		return Result{}, ctx.Err()
	}

	select {
	case <-req.done:
	case <-ctx.Done():
		return Result{}, ctx.Err()
	}

	return Result{Remaining: int64(cap(item.bucket) - len(item.bucket)), Limit: limit.Limit}, nil
}

func startLeaker(ctx context.Context, bucket <-chan request, limit Limit) {
	go func() {
		rate := float64(limit.Limit) / float64(limit.Unit)
		ticker := time.NewTicker(time.Duration(math.Floor(rate)))
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				select {
				case r := <-bucket:
					close(r.done)
				case <-ctx.Done():
					return
				default:
				}
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (l *LeakyBucketRateLimiter) SetLimits(limits map[string]Limit) {
	l.limits.Store(limits)
}

func (l *LeakyBucketRateLimiter) SetDefaultLimit(defaultLimit Limit) {
	l.defaultLimit.Store(defaultLimit)
}

func (l *LeakyBucketRateLimiter) SetTTL(ttl time.Duration) {
	l.ttl.Store(ttl)
}

func (l *LeakyBucketRateLimiter) cleanup(ctx context.Context) {
	ticker := time.NewTicker(l.cleanupInterval)

	for {
		select {
		case <-ticker.C:
			l.m.Lock()
			now := time.Now()
			for key, item := range l.state {
				if now.Sub(item.lastAccess) > l.ttl.Load().(time.Duration) {
					item.cancel()
					delete(l.state, key)
				}
			}
			l.m.Unlock()
		case <-ctx.Done():
			return
		}
	}
}
