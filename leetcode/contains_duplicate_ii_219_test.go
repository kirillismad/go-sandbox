package leetcode

import (
	"strconv"
	"testing"
)

func Test_containsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want bool
	}{
		{
			nums: []int{1, 2, 3, 1},
			k:    3,
			want: true,
		},
		{
			nums: []int{1, 0, 1, 1},
			k:    1,
			want: true,
		},
		{
			nums: []int{1, 2, 3, 1, 2, 3},
			k:    2,
			want: false,
		},
	}
	for i, c := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := containsNearbyDuplicate(c.nums, c.k); got != c.want {
				t.Errorf("got: %v, want: %v", got, c.want)
			}
		})
	}
}
