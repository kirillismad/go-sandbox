package cache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFIFOCache(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	cache := NewFIFOCache[int](2)
	cache.Put("a", 1)
	cache.Put("b", 2)

	val, found := cache.Get("b")
	r.True(found)
	r.Equal(2, val, "expected 2, got %v", val)

	cache.Put("c", 3)

	_, found = cache.Get("a")
	r.False(found)

	val, found = cache.Get("c")
	r.True(found)
	r.Equal(3, val, "expected 3, got %v", val)
}
