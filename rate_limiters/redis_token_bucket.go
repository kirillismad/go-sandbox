package rate_limiters

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisTokenBucketRateLimiter struct {
	rdb          *redis.Client
	limits       atomic.Value
	defaultLimit atomic.Value
	ttl          atomic.Value
	scriptHash   string
	m            sync.Mutex
	version      atomic.Int64
}

func NewRedisTokenBucketRateLimiter(rdb *redis.Client, limits map[string]Limit, opts ...Option[*RedisTokenBucketRateLimiter]) *RedisTokenBucketRateLimiter {
	l := &RedisTokenBucketRateLimiter{
		rdb: rdb,
	}
	l.SetLimits(limits)
	l.SetDefaultLimit(defaultLimit)
	l.SetTTL(defaultTTL)

	for _, o := range opts {
		o(l)
	}

	return l
}

func (l *RedisTokenBucketRateLimiter) Acquire(ctx context.Context, operation, ip string) (Result, error) {
	limitMap := l.limits.Load().(map[string]Limit)

	limit, found := limitMap[operation]
	if !found {
		limit = l.defaultLimit.Load().(Limit)
	}
	now := time.Now()
	key := fmt.Sprintf("rate_limit:%s:%s", operation, ip)

	keys := []string{key}
	args := []interface{}{now.Unix(), limit.Unit.Seconds(), limit.Limit, l.ttl.Load().(time.Duration).Seconds()}

	res, err := l.EvalSha(ctx, keys, args...)
	if err != nil {
		return Result{}, err
	}

	result := res.([]interface{})

	if result[0].(int64) == 0 {
		retryAfter := time.Unix(result[1].(int64), 0)
		return Result{}, NewErrRetryAfter(retryAfter)
	}

	return Result{
		Remaining: result[1].(int64),
		Limit:     limit.Limit,
	}, nil
}

func (l *RedisTokenBucketRateLimiter) SetLimits(limits map[string]Limit) {
	l.limits.Store(limits)
}

func (l *RedisTokenBucketRateLimiter) SetDefaultLimit(defaultLimit Limit) {
	l.defaultLimit.Store(defaultLimit)
}

func (l *RedisTokenBucketRateLimiter) SetTTL(ttl time.Duration) {
	l.ttl.Store(ttl)
}

func (l *RedisTokenBucketRateLimiter) uploadLua(ctx context.Context) (string, error) {
	content, err := os.ReadFile("./token_buckets.lua")
	if err != nil {
		return "", fmt.Errorf("ReadFile: %w", err)
	}

	res, err := l.rdb.ScriptLoad(ctx, string(content)).Result()
	if err != nil {
		return "", fmt.Errorf("ScriptLoad: %w", err)
	}

	return res, nil
}

func (l *RedisTokenBucketRateLimiter) Prepare(ctx context.Context) error {
	l.m.Lock()
	defer l.m.Unlock()

	hash, err := l.uploadLua(ctx)
	if err != nil {
		return err
	}
	l.scriptHash = hash
	l.version.Add(1)
	return nil
}

func (l *RedisTokenBucketRateLimiter) EvalSha(ctx context.Context, keys []string, args ...interface{}) (interface{}, error) {
	for attempt := 0; attempt < 5; attempt++ {
		res, err := l.rdb.EvalSha(ctx, l.scriptHash, keys, args...).Result()
		if err == nil {
			return res, nil
		}

		if strings.Contains(err.Error(), "NOSCRIPT") {
			return nil, fmt.Errorf("EvalSha: %w", err)
		}

		v := l.version.Load()
		l.m.Lock()

		if l.version.Load() == v {
			hash, err := l.uploadLua(ctx)
			if err != nil {
				l.m.Unlock()
				return nil, fmt.Errorf("uploadLua: %w", err)
			}
			l.scriptHash = hash
			l.version.Add(1)
		}
		l.m.Unlock()
	}

	res, err := l.rdb.EvalSha(ctx, l.scriptHash, keys, args...).Result()
	if err != nil {
		return nil, fmt.Errorf("EvalSha: %w", err)
	}
	return res, nil
}
