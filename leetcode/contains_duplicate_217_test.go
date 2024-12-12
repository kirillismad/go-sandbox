package leetcode

import (
	"strconv"
	"testing"
)

func Test_containsDuplicate(t *testing.T) {
	cases := []struct {
		nums []int
		want bool
	}{
		{
			nums: []int{1, 2, 3, 1},
			want: true,
		},
		{
			nums: []int{1, 2, 3, 4},
			want: false,
		},
		{
			nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			want: true,
		},
	}
	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := containsDuplicate(c.nums); got != c.want {
				t.Errorf("got: %v want: %v", got, c.want)
			}
		})
	}
}
