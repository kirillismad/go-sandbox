package leetcode

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_threeSum(t *testing.T) {
	cases := []struct {
		nums []int
		want [][]int
	}{
		{
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := threeSum(c.nums); !reflect.DeepEqual(got, c.want) {
				t.Errorf("got: %v, want: %v", got, c.want)
			}
		})
	}
}
