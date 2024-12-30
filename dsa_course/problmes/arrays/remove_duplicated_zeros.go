package arrays

// RemoveDuplicatedZeros is a function that removes duplicated zeros from the array.
// Complexity is O(n)
// Memory complexity is O(1)

// RemoveDuplicatedZeros([]int{1, 0, 0, 1, 1}) returns []int{1, 0, 1, 1}
func RemoveDuplicatedZeros(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	a := 1

	for i := 1; i < len(arr); i++ {
		if arr[i] == 0 {
			if arr[a-1] != 0 {
				arr[a] = arr[i]
				a++
			}
		} else {
			arr[a] = arr[i]
			a++
		}
	}
	return arr[:a]
}
