package leetcode

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

func removeDuplicates1(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	insertIndex := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[insertIndex-1] {
			nums[insertIndex] = nums[i]
			insertIndex++
		}
	}
	return insertIndex
}
