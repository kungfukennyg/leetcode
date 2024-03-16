package container

import (
	"fmt"
	"testing"
)

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{[]int{2, 3, 4, 5, 18, 17, 6}},
			want: 17,
		},
		{
			args: args{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			want: 49,
		},
		{
			args: args{[]int{1, 2, 1}},
			want: 2,
		},
		{
			args: args{[]int{1, 1}},
			want: 1,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Log("\n" + toString(tt.args.height))
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
