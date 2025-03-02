package tasks

func max(numbers []int, left int, right int) int {
	if len(numbers) == 0 {
		panic("empty slice")
	}

	maxIndex := left
	maxValue := numbers[maxIndex]

	for i := left; i <= right; i++ {
		if numbers[i] > maxValue {
			maxIndex = i
			maxValue = numbers[i]
		}
	}
	return maxIndex
}
