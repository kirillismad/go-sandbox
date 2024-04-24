package leetcode

import (
	"strconv"
	"testing"
)

func Test_minSubArrayLen(t *testing.T) {
	cases := []struct {
		target int
		nums   []int
		want   int
	}{
		{
			target: 7,
			nums:   []int{2, 3, 1, 2, 4, 3},
			want:   2,
		},
		{
			target: 4,
			nums:   []int{1, 4, 4},
			want:   1,
		},
		{
			target: 11,
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			want:   0,
		},
	}
	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := minSubArrayLen(c.target, c.nums); got != c.want {
				t.Errorf("got: %v, want: %v", got, c.want)
			}
		})
	}
}
