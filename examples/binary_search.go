package examples

import (
	"fmt"
	"sort"
)

func DemoBinarySearch() {
	demoBinarySearch()
}

func find(slice []int, value int) int {
	left, right := 0, len(slice)-1
	for left <= right {
		mid := left + (right-left)/2
		if slice[mid] < value {
			left = mid + 1
		} else if slice[mid] > value {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func findAll(slice []int, value int) []int {
	index := find(slice, value)
	result := make([]int, 0)

	if index != -1 {
		for i := index - 1; i >= 0 && slice[i] == slice[index]; i-- {
			result = append([]int{i}, result...)
		}
		for i := index; i < len(slice) && slice[i] == slice[index]; i++ {
			result = append(result, i)
		}

	}

	return result
}

func demoBinarySearch() {
	const LEN = 20
	s := make([]int, LEN)
	for i := range s {
		s[i] = rint(5)
	}
	sort.Ints(s)

	index := rint(LEN - 1)
	value := s[index]
	fmt.Printf("INDEX: %v, VALUE: %v\n", index, value)

	fmt.Println("SLICE:", s)
	fmt.Println("RESULT:", findAll(s, value))
	i, found := sort.Find(len(s), func(i int) int {
		if value < s[i] {
			return -1
		} else if value > s[i] {
			return 1
		} else {
			return 0
		}
	})
	fmt.Printf("sort.Find -- INDEX: %v, FOUND: %v\n", i, found)
}
