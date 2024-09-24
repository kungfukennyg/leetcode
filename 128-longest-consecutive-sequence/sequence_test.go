package sequence

import (
	"strconv"
	"testing"
)

func Test_longestConsecutive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				nums: []int{1, 2, 0, 1},
			},
			want: 3,
		},
		{
			args: args{
				nums: []int{0, -1},
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{},
			},
			want: 0,
		},
		{
			args: args{
				nums: []int{100, 4, 200, 1, 3, 2},
			},
			want: 4,
		},
		{
			args: args{
				nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			},
			want: 9,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := longestConsecutive(tt.args.nums); got != tt.want {
				t.Errorf("longestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
