package cache

import (
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestLRUCache(t *testing.T) {
	t.Parallel()

	t.Run("default", func(t *testing.T) {
		r := require.New(t)

		cache := NewLRUCache[int](3)

		for i := range 3 {
			cache.Put(strconv.Itoa(i), i)
		}

		for i := range 3 {
			res, ok := cache.Get(strconv.Itoa(i))
			r.True(ok)
			r.Equal(i, res)
		}

		cache.Put("3", 3)
		_, ok := cache.Get("0")
		r.False(ok)

		res, ok := cache.Get("3")
		r.True(ok)
		r.Equal(3, res)
	})
}
