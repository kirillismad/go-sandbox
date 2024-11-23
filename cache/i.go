package cache

type Cache[T any] interface {
	Get(key string) (T, bool)
	Put(key string, value T)
}

// Entry Cache entry
type Entry[T any] struct {
	key   string
	value T
}
