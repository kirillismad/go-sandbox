package leetcode

import (
	"strconv"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	cases := []struct {
		s    string
		want bool
	}{
		{
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			s:    "race a car",
			want: false,
		},
		{
			s:    "0P",
			want: false,
		},
	}
	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := isPalindrome(c.s); got != c.want {
				t.Errorf("got: %v, want: %v", got, c.want)
			}
		})
	}
}
