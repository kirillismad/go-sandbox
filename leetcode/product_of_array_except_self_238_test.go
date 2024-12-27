package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_productExceptSelf(t *testing.T) {
	t.Parallel()

	tests := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{2, 3, 4, 5},
			expected: []int{60, 40, 30, 24},
		},
		{
			nums:     []int{1, 2, 3, 4, 5},
			expected: []int{120, 60, 40, 30, 24},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			res := productExceptSelf(tt.nums)
			require.ElementsMatch(t, tt.expected, res)
		})
	}
}
