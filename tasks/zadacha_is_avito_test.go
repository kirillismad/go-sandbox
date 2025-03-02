package tasks

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMax(t *testing.T) {
	tests := []struct {
		numbers []int
		left    int
		right   int
		want    int
	}{
		{[]int{1, 3, 1, 2}, 0, 3, 1},
		{[]int{1, 3, 1, 2}, 1, 3, 1},
		{[]int{1, 3, 1, 2}, 2, 3, 3},
		{[]int{1, 3, 1, 2}, 0, 2, 1},
		{[]int{5, 4, 3, 2, 1}, 0, 4, 0},
		{[]int{1, 2, 3, 4, 5}, 0, 4, 4},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			got := max(tt.numbers, tt.left, tt.right)
			require.Equal(t, tt.want, got, "max(%v, %d, %d)", tt.numbers, tt.left, tt.right)
		})
	}

	t.Run("empty slice", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("max did not panic on empty slice")
			}
		}()
		max([]int{}, 0, 0)
	})

	t.Run("empty slice", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("max did not panic on empty slice")
			}
		}()
		max([]int{}, 0, 0)
	})
}
