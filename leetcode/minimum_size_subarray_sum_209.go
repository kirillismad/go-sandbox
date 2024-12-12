package leetcode

import "math"

// https://leetcode.com/problems/minimum-size-subarray-sum/

// #sliding-window

func minSubArrayLen(target int, nums []int) int {
	var (
		sum       = 0
		minLen    = math.MaxInt
		leftIndex = 0
	)

	for rightIndex := 0; rightIndex < len(nums); rightIndex++ {
		sum += nums[rightIndex]

		for sum >= target {
			minLen = min(minLen, rightIndex-leftIndex+1) // totalLen = i - leftIndex + 1
			sum -= nums[leftIndex]
			leftIndex++
		}
	}
	if minLen != math.MaxInt {
		return minLen
	}
	return 0
}
