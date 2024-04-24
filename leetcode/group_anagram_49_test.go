package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	cases := []struct {
		strs []string
		want [][]string
	}{
		{
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := groupAnagrams(c.strs)
			fmt.Printf("got: %v\n", got)
		})
	}
}
