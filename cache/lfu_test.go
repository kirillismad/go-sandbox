package cache

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLFUCache(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	cache := NewLFUCache[int](2)

	cache.Put("a", 1)
	cache.Put("b", 2)

	val, found := cache.Get("a")
	r.True(found)
	r.Equal(1, val, "expected 1, got %v", val)

	cache.Put("c", 3)

	_, found = cache.Get("b")
	r.False(found, "expected b to be evicted")

	val, found = cache.Get("c")
	r.True(found)
	r.Equal(3, val, "expected 3, got %v", val)
}
