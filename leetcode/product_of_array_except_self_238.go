package leetcode

// https://leetcode.com/problems/product-of-array-except-self/
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = 1
	}
	prevLeft, prevRight := 1, 1
	for i := 0; i < len(nums)-1; i++ {
		j := len(nums) - i - 1

		prevLeft = prevLeft * nums[i]
		prevRight = prevRight * nums[j]

		result[i+1] *= prevLeft
		result[j-1] *= prevRight
	}

	return result
}
