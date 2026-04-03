package std

import "iter"

func Reversed[E any, S ~[]E](slice S) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for i := len(slice) - 1; i >= 0; i-- {
			if !yield(i, slice[i]) {
				return
			}
		}
	}
}
