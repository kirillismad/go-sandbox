package cache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClockCache(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	cache := NewClockCache[int](2)

	// Test inserting and retrieving values
	cache.Put("a", 1)
	cache.Put("b", 2)

	val, found := cache.Get("a")
	r.True(found)
	r.Equal(1, val, "expected 1, got %v", val)

	val, found = cache.Get("b")
	r.True(found)
	r.Equal(2, val, "expected 2, got %v", val)

	// Test eviction
	cache.Put("c", 3)
	_, found = cache.Get("a")
	r.False(found, "expected key 'a' to be evicted")

	val, found = cache.Get("c")
	r.True(found)
	r.Equal(3, val, "expected 3, got %v", val)
}
