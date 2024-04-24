package leetcode

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name     string
		args     args
		want     int
		wantNums []int
	}{
		{
			name:     "1",
			args:     args{nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3}},
			want:     7,
			wantNums: []int{0, 0, 1, 1, 2, 3, 3},
		},
		{
			name:     "2",
			args:     args{nums: []int{1, 1, 1, 2, 2, 3}},
			want:     5,
			wantNums: []int{1, 1, 2, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates2(tt.args.nums); got != tt.want {
				t.Errorf("%v != %v", got, tt.want)
			}
			for i := 0; i < tt.want; i++ {
				if tt.args.nums[i] != tt.wantNums[i] {
					t.Errorf("%v != %v", tt.args.nums[i], tt.wantNums[i])
				}
			}
		})
	}
}
