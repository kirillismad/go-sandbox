package leetcode

import (
	"slices"
	"testing"
)

func Test_removeDuplicates1(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	want := []int{0, 1, 2, 3, 4}
	idx := removeDuplicates1(nums)

	if slices.Compare(want, nums[:idx]) != 0 {
		t.Errorf("got: %v, want: %v", nums[:idx], want)
	}
}
