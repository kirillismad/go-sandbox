package leetcode

import "sort"

// https://leetcode.com/problems/longest-consecutive-sequence/

func longestConsecutive(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	sort.Ints(nums)

	ans := 0
	counter := 1
	for i := 1; i < len(nums); i++ {
		switch nums[i] - nums[i-1] {
		case 0:
			continue
		case 1:
			counter++
		default:
			ans = max(counter, ans)
			counter = 1
		}
	}

	return max(counter, ans)
}
