package tasks

import (
	"reflect"
	"testing"
)

func TestAddTwoSlices(t *testing.T) {
	cases := []struct {
		name string
		s1   []int
		s2   []int
		want []int
	}{
		{
			name: "7899+999",
			s1:   []int{7, 8, 9, 9},
			s2:   []int{9, 9, 9},
			want: []int{8, 8, 9, 8},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := AddTwoSlices(c.s1, c.s2); !reflect.DeepEqual(got, c.want) {
				t.Errorf("AddTwoSlices() = %v, want %v", got, c.want)
			}
		})
	}
}
