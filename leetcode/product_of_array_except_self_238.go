package leetcode

// https://leetcode.com/problems/product-of-array-except-self/
func productExceptSelf(nums []int) []int {
	totalLen := len(nums)
	result := make([]int, totalLen)
	for i := 0; i < totalLen; i++ {
		result[i] = 1
	}
	prevLeft, prevRight := 1, 1
	for i := 0; i < totalLen; i++ {
		j := totalLen - i - 1
		result[i] *= prevLeft
		result[j] *= prevRight

		prevLeft = prevLeft * nums[i]
		prevRight = prevRight * nums[j]
	}

	return result
}
