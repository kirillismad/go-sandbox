package cache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFIFOCache_PutAndGet(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	cache := NewFIFOShardsCache[int](2, 2)
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)
	cache.Put("d", 4)

	val, found := cache.Get("a")
	r.True(found)
	r.Equal(1, val, "expected 1, got %v", val)

	val, found = cache.Get("b")
	r.True(found)
	r.Equal(2, val, "expected 2, got %v", val)

	val, found = cache.Get("c")
	r.True(found)
	r.Equal(3, val, "expected 3, got %v", val)

	val, found = cache.Get("d")
	r.True(found)
	r.Equal(4, val, "expected 4, got %v", val)

	cache.Put("e", 5)

	_, found = cache.Get("a")
	r.False(found)

	val, found = cache.Get("e")
	r.True(found)
	r.Equal(5, val, "expected 5, got %v", val)
}
