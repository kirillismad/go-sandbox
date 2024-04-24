package utils

func Must[T any](x T, e error) T {
	if e != nil {
		panic(e)
	}
	return x
}

func SliceMap[T any, R any](s []T, f func(x T) R) []R {
	result := make([]R, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}
