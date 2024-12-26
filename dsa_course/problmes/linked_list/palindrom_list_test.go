package linkedlist

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPalindrome(t *testing.T) {
	t.Parallel()

	tests := []struct {
		head *ListNode
		want bool
	}{
		{
			head: NewList([]int{1, 2, 2, 1}),
			want: true,
		},
		{
			head: NewList([]int{1, 2, 3, 2, 1}),
			want: true,
		},
		{
			head: NewList([]int{1, 2, 3, 4, 1}),
			want: false,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			res := isPalindrome(tt.head)
			require.Equal(t, tt.want, res)
		})
	}
}
