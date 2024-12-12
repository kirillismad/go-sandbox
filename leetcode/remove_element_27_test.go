package leetcode

import (
	"slices"
	"testing"
)

func Test_removeElement(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	want := []int{0, 1, 3, 0, 4}

	idx := removeElement(nums, val)

	if slices.Compare(want, nums[:idx]) != 0 {
		t.Errorf("got: %v, want: %v", nums[:idx], want)
	}
}
