package leetcode

// https://leetcode.com/problems/merge-sorted-array/

func merge(nums1 []int, m int, nums2 []int, n int) {
	// при n == 0, массивы смержены
	for n != 0 {
		// nums1[m-1] == последний элемент в nums1
		// nums2[n-1] == последний элемент в nums2
		// Если m == 0 тогда все элементы nums1 были перемещены в хвост
		// наприер nums1 = [4,5,6,0,0,0], nums2 = [1,2,3], тогда при m == 0 nums1 == [4,5,6,4,5,6]
		// на место nums[0], nums[1], nums[2] придут [1,2,3] и будет [1,2,3,4,5,6]
		if m != 0 && nums1[m-1] > nums2[n-1] {
			// nums1[m+n-1] == last index in result array on current iteration
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}
