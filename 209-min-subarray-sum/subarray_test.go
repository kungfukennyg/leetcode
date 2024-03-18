package subarray

import (
	"fmt"
	"testing"
)

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				target: 6,
				nums:   []int{10, 2, 3},
			},
			want: 1,
		},
		{
			args: args{
				target: 7,
				nums:   []int{8},
			},
			want: 1,
		},
		{
			args: args{
				target: 15,
				nums:   []int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8},
			},
			want: 2,
		},
		{
			args: args{
				target: 15,
				nums:   []int{1, 2, 3, 4, 5},
			},
			want: 5,
		},
		{
			args: args{
				target: 213,
				nums:   []int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12},
			},
			want: 8,
		},
		{
			args: args{
				target: 7,
				nums:   []int{2, 3, 1, 2, 4, 3},
			},
			want: 2,
		},
		{
			args: args{
				target: 4,
				nums:   []int{1, 4, 4},
			},
			want: 1,
		},
		{
			args: args{
				target: 11,
				nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			},
			want: 0,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := minSubArrayLen(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
