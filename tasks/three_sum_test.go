package tasks

import (
	"bufio"
	"os"
	"sandbox/utils"
	"strconv"
	"strings"
	"testing"
)

func twoSumsFile(filename string) []int {
	file := utils.Must(os.Open(filename))
	defer file.Close()

	result := make([]int, 0, 1024)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i := utils.Must(strconv.Atoi(strings.Trim(scanner.Text(), " ")))
		result = append(result, i)
	}

	return result
}
func TestThreeSum(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "1k",
			nums: twoSumsFile("./testdata/1k.txt"),
			want: 70,
		},
		{
			name: "2k",
			nums: twoSumsFile("./testdata/2k.txt"),
			want: 528,
		},
		{
			name: "4k",
			nums: twoSumsFile("./testdata/4k.txt"),
			want: 4039,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := ThreeSum(c.nums); got != c.want {
				t.Errorf("ThreeSum() = %v, want %v", got, c.want)
			}
		})
	}
}

func TestThreeSumFast(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "1k",
			nums: twoSumsFile("./testdata/1k.txt"),
			want: 70,
		},
		{
			name: "2k",
			nums: twoSumsFile("./testdata/2k.txt"),
			want: 528,
		},
		{
			name: "4k",
			nums: twoSumsFile("./testdata/4k.txt"),
			want: 4039,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := ThreeSumFast(c.nums); got != c.want {
				t.Errorf("ThreeSumFast() = %v, want %v", got, c.want)
			}
		})
	}
}
