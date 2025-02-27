package rate_limiters

import (
	"context"
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

type request struct {
	done chan struct{}
}

type CapLimit struct {
	Limit
	Cap int64
}

type lbItem struct {
	bucket     chan request
	cancel     context.CancelFunc
	lastAccess time.Time
}

var defaultCapLimit = CapLimit{
	Limit: defaultLimit,
	Cap:   64,
}

// LeakyBucketRateLimiter implements a rate limiter using the leaky bucket algorithm.
// It maintains a state map to track the rate limit information for different keys.
// The structure uses a mutex for synchronization and atomic values for limits, default limit, and TTL.
// The cleanupInterval specifies the duration for periodic cleanup of expired entries.
type LeakyBucketRateLimiter struct {
	state           map[string]lbItem
	mu              sync.Mutex
	limits          atomic.Value
	defaultLimit    atomic.Value
	ttl             atomic.Value
	cleanupInterval time.Duration
}

// NewLeakyBucketRateLimiter creates a new instance of LeakyBucketRateLimiter with the provided context, limits, and options.
// It initializes the rate limiter with default values for TTL, default limit, and cleanup interval, which can be overridden
// by the provided options. The cleanup process is started in a separate goroutine.
//
// Parameters:
//   - ctx: The context to control the lifecycle of the rate limiter.
//   - limits: A map of string keys to CapLimit values defining the rate limits for different keys.
//   - opts: A variadic list of Option functions to customize the rate limiter.
//
// Returns:
//   - A pointer to the newly created LeakyBucketRateLimiter.
func NewLeakyBucketRateLimiter(ctx context.Context, limits map[string]CapLimit, opts ...Option[*LeakyBucketRateLimiter]) *LeakyBucketRateLimiter {
	l := &LeakyBucketRateLimiter{
		state:           make(map[string]lbItem),
		cleanupInterval: defaultCleanupInterval,
	}

	l.SetLimits(limits)

	l.SetTTL(defaultTTL)
	l.SetDefaultLimit(defaultCapLimit)
	l.SetCleanupInterval(defaultCleanupInterval)

	for _, opt := range opts {
		opt(l)
	}

	go l.cleanup(ctx)

	return l
}

func (l *LeakyBucketRateLimiter) Acquire(ctx context.Context, operation, ip string) (Result, error) {
	limitMap := l.limits.Load().(map[string]CapLimit)

	limit, found := limitMap[operation]
	if !found {
		limit = l.defaultLimit.Load().(CapLimit)
	}

	key := fmt.Sprintf("rate_limit:%s:%s", operation, ip)

	l.mu.Lock()
	item, found := l.state[key]
	if !found {
		newCtx, cancel := context.WithCancel(context.Background())
		item = lbItem{bucket: make(chan request, limit.Cap), cancel: cancel}

		startLeaker(newCtx, item.bucket, limit)
		l.state[key] = item
	}
	l.mu.Unlock()

	if len(item.bucket) >= int(limit.Cap) {
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

	return Result{Remaining: int64(cap(item.bucket) - len(item.bucket)), Limit: limit.Cap}, nil
}

func startLeaker(ctx context.Context, bucket <-chan request, limit CapLimit) {
	go func() {
		rate := float64(limit.Limit.Unit) / float64(limit.Limit.Limit)

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

func (l *LeakyBucketRateLimiter) SetLimits(limits map[string]CapLimit) {
	l.limits.Store(limits)
}

func (l *LeakyBucketRateLimiter) SetDefaultLimit(defaultLimit CapLimit) {
	l.defaultLimit.Store(defaultLimit)
}

func (l *LeakyBucketRateLimiter) SetTTL(ttl time.Duration) {
	l.ttl.Store(ttl)
}

func (l *LeakyBucketRateLimiter) SetCleanupInterval(interval time.Duration) {
	l.cleanupInterval = interval
}

func (l *LeakyBucketRateLimiter) cleanup(ctx context.Context) {
	ticker := time.NewTicker(l.cleanupInterval)

	for {
		select {
		case <-ticker.C:
			l.mu.Lock()
			now := time.Now()
			for key, item := range l.state {
				if now.Sub(item.lastAccess) > l.ttl.Load().(time.Duration) {
					item.cancel()
					delete(l.state, key)
				}
			}
			l.mu.Unlock()
		case <-ctx.Done():
			return
		}
	}
}
