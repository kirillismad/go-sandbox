package leetcode

import (
	"strconv"
	"testing"
)

func Test_findMaxAverage(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
		want float64
	}{
		{
			nums: []int{0, 1, 2, 3, 4, 5, 6, 7},
			k:    2,
			want: 6.5,
		},
		{
			nums: []int{1, 12, -5, -6, 50, 3},
			k:    4,
			want: 12.75000,
		},
	}
	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := findMaxAverage(c.nums, c.k); got != c.want {
				t.Errorf("got: %v, want: %v", got, c.want)
			}
		})
	}
}
