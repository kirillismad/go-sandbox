package leetcode

// https://leetcode.com/problems/product-of-array-except-self/
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = 1
	}
	prevLeft, prevRight := 1, 1
	for i := 0; i < len(nums); i++ {
		j := len(nums) - i - 1

		result[i] *= prevLeft
		result[j] *= prevRight

		prevLeft = prevLeft * nums[i]
		prevRight = prevRight * nums[j]
	}

	return result
}
