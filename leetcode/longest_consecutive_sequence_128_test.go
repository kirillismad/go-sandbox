package leetcode

import (
	"math/rand"
	"strconv"
	"testing"
)

const bigValue = 1_000_000

func alotOfInts() []int {
	result := make([]int, bigValue)
	for i := 0; i < bigValue; i++ {
		result[i] = i
	}
	rand.Shuffle(len(result), func(i, j int) { result[i], result[j] = result[j], result[i] })
	return result
}

func Test_longestConsecutive(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
		{
			nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want: 9,
		},
		{
			nums: alotOfInts(),
			want: bigValue,
		},
	}
	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := longestConsecutive(c.nums); got != c.want {
				t.Errorf("longestConsecutive() = %v, want %v", got, c.want)
			}
		})
	}
}
