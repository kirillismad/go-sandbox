package rate_limiters

import "time"

type Option[T any] func(l T)

type ISetDefaultLimit[L any] interface {
	SetDefaultLimit(l L)
}

func WithDefaultLimit[T ISetDefaultLimit[L], L any](limit L) Option[T] {
	return func(l T) {
		l.SetDefaultLimit(limit)
	}
}

type ISetTTL interface {
	SetTTL(ttl time.Duration)
}

func WithTTL[T ISetTTL](ttl time.Duration) Option[T] {
	return func(l T) {
		l.SetTTL(ttl)
	}
}

type ISetCleanupInterval interface {
	SetCleanupInterval(cleanupInterval time.Duration)
}

func WithCleanupInterval[T ISetCleanupInterval](cleanupInterval time.Duration) Option[T] {
	return func(l T) {
		l.SetCleanupInterval(cleanupInterval)
	}
}

var defaultLimit = Limit{
	Unit:  time.Minute,
	Limit: 1,
}

const defaultTTL = 1 * time.Hour
const defaultCleanupInterval = 1 * time.Minute
