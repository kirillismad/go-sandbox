package leetcode

import "testing"

func Test_longestValidParentheses(t *testing.T) {
	cases := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "1",
			s:    "(()",
			want: 2,
		},
		{
			name: "2",
			s:    ")()())",
			want: 4,
		},
		{
			name: "3",
			s:    "",
			want: 0,
		},
		{
			name: "4",
			s:    "()(()",
			want: 2,
		},
		{
			name: "5",
			s:    "()(())",
			want: 6,
		},
		{
			name: "6",
			s:    "))(())",
			want: 4,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := longestValidParentheses(c.s); got != c.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, c.want)
			}
		})
	}
}
