package leetcode

import "testing"

func Test_isValid(t *testing.T) {

	cases := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "1",
			s:    "({[]})",
			want: true,
		},
		{
			name: "2",
			s:    "(}",
			want: false,
		},
		{
			name: "3",
			s:    "(",
			want: false,
		},
		{
			name: "4",
			s:    ")",
			want: false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := isValid(c.s); got != c.want {
				t.Errorf("isValid() = %v, want %v", got, c.want)
			}
		})
	}
}
