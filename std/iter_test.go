package std

import (
	"testing"

	"github.com/brianvoe/gofakeit/v7"
)

func TestIter(t *testing.T) {
	genSlice := func(n int) []int {
		result := make([]int, n)
		for i := 0; i < n; i++ {
			result[i] = gofakeit.Int()
		}
		return result
	}
	t.Run("reverse", func(t *testing.T) {
		in := genSlice(100)

		iter := Reversed(in)

		for idx, val := range iter {
			expectedVal := in[idx]
			if val != expectedVal {
				t.Errorf("Expected value %d at index %d, got %d", expectedVal, idx, val)
			}
		}

	})
}
