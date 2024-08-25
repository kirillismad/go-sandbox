package sort

import "cmp"

type QuickSorter[T cmp.Ordered] struct {
}

func NewQuickSorter[T cmp.Ordered]() *QuickSorter[T] {
	return &QuickSorter[T]{}
}

func (s *QuickSorter[T]) Sort(items []T) {
	s.sort(items, 0, len(items)-1)
}

func (s *QuickSorter[T]) sort(slice []T, left int, right int) {
	if left < right {
		p := s.partition(slice, left, right)
		s.sort(slice, left, p)
		s.sort(slice, p+1, right)
	}
}

func (s *QuickSorter[T]) partition(slice []T, left int, right int) int {
	pivot := slice[(left+right)/2]
	for {
		for slice[left] < pivot {
			left++
		}
		for slice[right] > pivot {
			right--
		}
		if left >= right {
			return right
		}
		slice[left], slice[right] = slice[right], slice[left]
	}
}
