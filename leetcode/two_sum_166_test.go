package leetcode

import (
	"slices"
	"strconv"
	"testing"
)

func Test_twoSum1(t *testing.T) {
	cases := []struct {
		numbers []int
		target  int
		want    []int
	}{
		{
			numbers: []int{2, 7, 11, 15},
			target:  9,
			want:    []int{1, 2},
		},
		{
			numbers: []int{2, 3, 4},
			target:  6,
			want:    []int{1, 3},
		},
		{
			numbers: []int{-1, 0},
			target:  -1,
			want:    []int{1, 2},
		},
	}
	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := twoSum1(c.numbers, c.target); slices.Compare(got, c.want) != 0 {
				t.Errorf("twoSum1() = %v, want %v", got, c.want)
			}
		})
	}
}
