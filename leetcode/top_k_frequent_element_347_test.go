package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_topKFrequent(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 2, 2, 3}
	k := 2
	result := topKFrequent(nums, k)

	require.Equal(t, []int{2, 1}, result)
}
