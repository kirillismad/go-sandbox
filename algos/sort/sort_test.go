package sort

import (
	"math/rand"
	gosort "sort"
	"testing"
)

func TestSorter(t *testing.T) {
	nums1 := rand.Perm(2056)
	nums2 := make([]int, len(nums1))
	copy(nums2, nums1)

	cases := []struct {
		name   string
		sorter Sorter[int]
		items  []int
	}{
		{
			name:   "Select",
			sorter: NewSelectSorter[int](),
			items:  nums1,
		},
		{
			name:   "Insert",
			sorter: NewInsertSorter[int](),
			items:  nums2,
		},
		{
			name:   "Shell",
			sorter: NewShellSorter[int](),
			items:  nums2,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.sorter.Sort(c.items)
			isSorted := gosort.IsSorted(gosort.IntSlice(c.items))
			if !isSorted {
				t.Errorf("%v is not sorted", c.items)
			}
		})
	}
}
