package jump

import (
	"strconv"
	"testing"
)

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{nums: []int{1, 1, 1, 0}},
			want: true,
		}, {
			args: args{nums: []int{2, 3, 1, 1, 4}},
			want: true,
		},
		{
			args: args{nums: []int{3, 2, 1, 0, 4}},
			want: false,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
