package leetcode

import "math"

// https://leetcode.com/problems/maximum-average-subarray-i/

// #sliding-window

func findMaxAverage(nums []int, k int) float64 {
	var (
		leftValue = 0.0
		sum       = 0.0
		ans       = math.Inf(-1)
	)

	for rightIndex := 0; rightIndex < len(nums); rightIndex++ {
		sum += float64(nums[rightIndex])

		if rightIndex+1 >= k {
			sum -= leftValue

			ans = max(sum/float64(k), ans)

			leftValue = float64(nums[rightIndex+1-k]) // leftIndex = rightIndex + 1 - k
		}
	}
	return ans
}
