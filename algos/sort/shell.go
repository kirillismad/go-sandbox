package sort

import (
	"cmp"
)

type ShellSorter[T cmp.Ordered] struct{}

func NewShellSorter[T cmp.Ordered]() ShellSorter[T] {
	return ShellSorter[T]{}
}

func (s ShellSorter[T]) Sort(items []T) {
	h := 1

	for h < len(items)/3 {
		h = h*3 + 1
	}
	for h > 0 {
		for i := h; i < len(items); i++ {
			for j := i; j >= h && items[j] < items[j-h]; j -= h {
				items[j], items[j-h] = items[j-h], items[j]
			}
		}
		h /= 3
	}
}
