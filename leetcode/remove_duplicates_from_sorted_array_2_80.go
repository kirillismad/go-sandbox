package leetcode

// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/

func removeDuplicates2(nums []int) int {
	// любой состав массива из 2 или менее элементов удовлетворяет задаче
	if len(nums) <= 2 {
		return len(nums)
	}
	insertIndex := 2
	for i := 2; i < len(nums); i++ {
		// если nums[i] == nums[insertIndex-2]
		// то в сущности это обозначает что nums[i] уже является третим одинаковым элементом
		// в таком случае не надо ничего делать кроме как пойти дальше в поисках отличного элемента
		// если nums[i] != nums[insertIndex-2]
		// то в сущности это означает, что отличный элемент был найден и его надо вставить по индексу insertIndex и ++
		if nums[i] != nums[insertIndex-2] {
			nums[insertIndex] = nums[i]
			insertIndex++
		}
	}
	return insertIndex
}
