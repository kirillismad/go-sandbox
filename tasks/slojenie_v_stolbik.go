package tasks

func AddTwoSlices(s1, s2 []int) []int {
	r := make([]int, 0)
	var sum int
	for i, j := len(s1)-1, len(s2)-1; i >= 0 || j >= 0; {
		if i >= 0 {
			sum += s1[i]
			i--
		}
		if j >= 0 {
			sum += s2[j]
			j--
		}

		r = append(r, sum%10)

		sum /= 10
	}
	if sum != 0 {
		r = append(r, sum)
	}
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}
