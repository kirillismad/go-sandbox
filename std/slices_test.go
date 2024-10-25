package std

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSlicesAppend(t *testing.T) {
	t.Parallel()

	durtyAppend := func(s []int) []int {
		s = append(s, 42)
		return s
	}
	t.Run("dirty func, nil slice", func(t *testing.T) {
		t.Parallel()

		r := require.New(t)

		var slice []int

		durtyAppend(slice)

		r.Nil(slice)
		t.Logf("slice = %v, len = %d, cap = %d", slice, len(slice), cap(slice))
	})

	t.Run("dirty func, slice literal", func(t *testing.T) {
		t.Parallel()

		r := require.New(t)

		var slice []int = []int{1, 2}

		durtyAppend(slice)

		r.ElementsMatch([]int{1, 2}, slice)
		t.Logf("slice = %v, len = %d, cap = %d", slice, len(slice), cap(slice))
	})

	t.Run("dirty func, slice make len = 2, cap = 2", func(t *testing.T) {
		t.Parallel()

		r := require.New(t)

		var slice []int = make([]int, 2)

		durtyAppend(slice)

		r.ElementsMatch([]int{0, 0}, slice)
		t.Logf("slice = %v, len = %d, cap = %d", slice, len(slice), cap(slice))
	})

	t.Run("dirty func, slice make len = 0, cap = 2", func(t *testing.T) {
		t.Parallel()

		r := require.New(t)

		var slice []int = make([]int, 0, 2)

		durtyAppend(slice)

		r.ElementsMatch([]int{}, slice)
		t.Logf("slice = %v, len = %d, cap = %d", slice, len(slice), cap(slice))
	})
}

func TestSlicesIndexing(t *testing.T) {
	t.Parallel()

	durtyChange := func(s []int) []int {
		s[0] = 42
		return s
	}

	t.Run("durty func, change index value", func(t *testing.T) {
		t.Parallel()

		r := require.New(t)

		slice := []int{1, 2}

		durtyChange(slice)

		r.ElementsMatch([]int{42, 2}, slice)
		t.Logf("slice = %v, len = %d, cap = %d", slice, len(slice), cap(slice))
	})
}

func TestSlicesPazzles(t *testing.T) {
	t.Parallel()

	t.Run("1", func(t *testing.T) {
		t.Parallel()
		r := require.New(t)

		a := [...]int{0, 1, 2, 3}
		x := a[:1]          // [0], len = 1, cap = 4
		y := a[2:]          // [2, 3], len = 2, cap = 2
		x = append(x, y...) // [0, 2, 3] len = 3, cap = 4, a = [0, 2, 3, 3], y = [3, 3]

		x = append(x, y...) // [0, 2, 3, 3, 3], len = 5, cap = 8

		r.ElementsMatch([...]int{0, 2, 3, 3}, a)
		r.ElementsMatch([]int{0, 2, 3, 3, 3}, x)
		r.ElementsMatch([]int{3, 3}, y)
	})

	t.Run("2", func(t *testing.T) {
		t.Parallel()
		r := require.New(t)

		a := []byte("ba") // [b, a] len = 2, cap = 2
		t.Logf("slice = %v, len = %d, cap = %d", a, len(a), cap(a))

		a1 := append(a, 'd')
		a2 := append(a, 'g')

		r.Equal("bad", string(a1))
		r.Equal("bag", string(a2))
	})
}
