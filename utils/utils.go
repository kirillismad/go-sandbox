package utils

func Must[T any](x T, e error) T {
	if e != nil {
		panic(e)
	}
	return x
}

func Map[S, D any](items []S, f func(S) D) []D {
	result := make([]D, len(items))
	for i, v := range items {
		result[i] = f(v)
	}
	return result
}
