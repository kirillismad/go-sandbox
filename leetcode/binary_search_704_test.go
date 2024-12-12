package leetcode

import (
	"strconv"
	"testing"
)

func Test_search(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			want:   4,
		},
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			want:   -1,
		},
	}
	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := search(c.nums, c.target); got != c.want {
				t.Errorf("got: %v, want: %v", got, c.want)
			}
		})
	}
}
