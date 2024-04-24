package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}

	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		// после того как первый элемент был обработан (i > 0)
		// при проверке i-го элемента проверяю что он является дубликатом
		// дубликаты пропускаю, этим достигается условие отсутствия одинаковых триплетов
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		// задача сводится к two_sum_166, где target == -nums[i]
		left := i + 1
		right := len(nums) - 1
		for left < right {
			switch sum := nums[i] + nums[left] + nums[right]; {
			case sum > 0:
				right--
			case sum < 0:
				left++
			default:
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				right--
				// пропускаю дубликаты с правой стороны
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}

	return ans
}
