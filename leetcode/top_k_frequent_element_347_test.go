package leetcode

import (
	"fmt"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 2, 2, 3}
	k := 2
	result := topKFrequent(nums, k)

	fmt.Printf("topKFrequent(%v, %v) = %v\n", nums, k, result)
}
