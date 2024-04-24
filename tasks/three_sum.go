package tasks

import (
	"sort"
)

func ThreeSum(nums []int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					cnt++
				}
			}
		}
	}
	return cnt
}

func binarySearch(s []int, q int) (int, bool) {
	if i := sort.Search(len(s), func(i int) bool { return s[i] >= q }); i < len(s) && s[i] == q {
		return i, true
	}
	return -1, false
}

func ThreeSumFast(nums []int) int {
	newNums := make([]int, len(nums))
	copy(newNums, nums)
	sort.Ints(newNums)

	cnt := 0
	for i := 0; i < len(newNums); i++ {
		for j := i + 1; j < len(newNums); j++ {
			if k, found := binarySearch(newNums, -(newNums[i] + newNums[j])); found && k > j {
				cnt++
			}
		}
	}
	return cnt
}
