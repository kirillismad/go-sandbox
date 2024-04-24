package leetcode

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
// #two-pointers

func twoSum1(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		switch {
		case numbers[left]+numbers[right] < target:
			left++
		case numbers[left]+numbers[right] > target:
			right--
		default:
			return []int{left + 1, right + 1}
		}
	}
	return nil
}
