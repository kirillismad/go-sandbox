package leetcode

import (
	"sort"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	cases := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{4, 5, 6},
				n:     3,
			},
		},
		{
			name: "2",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
		},
		{
			name: "3",
			args: args{
				nums1: []int{0},
				m:     0,
				nums2: []int{1},
				n:     1,
			},
		},
		{
			name: "4",
			args: args{
				nums1: []int{1},
				m:     1,
				nums2: []int{},
				n:     0,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			merge(c.args.nums1, c.args.m, c.args.nums2, c.args.n)
			result := sort.SliceIsSorted(c.args.nums1, func(i, j int) bool { return c.args.nums1[i] < c.args.nums1[j] })
			if !result {
				t.Error("invalid result")
			}
		})
	}
}
