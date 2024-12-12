package sort

import "cmp"

type Sorter[T cmp.Ordered] interface {
	Sort([]T)
}
