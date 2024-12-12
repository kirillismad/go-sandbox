package leetcode

import (
	"fmt"
	"testing"
)

func Test_twoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	indexes := twoSum(nums, 18)
	fmt.Printf("twoSum(%v, %v) = %v\n", nums, 9, indexes)
}
