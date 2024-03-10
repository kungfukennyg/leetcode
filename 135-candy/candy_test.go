package candy

import (
	"fmt"
	"testing"
)

func Test_candy(t *testing.T) {
	type args struct {
		ratings []int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{[]int{1, 6, 10, 8, 7, 3, 2}},
			want: 18,
		},
		{
			args: args{[]int{1, 3, 4, 5, 2}},
			want: 11,
		},
		{
			args: args{[]int{1, 2, 87, 87, 87, 2, 1}},
			want: 13,
		},
		{
			args: args{[]int{1, 0, 2}},
			want: 5,
		},
		{
			args: args{[]int{1, 2, 2}},
			want: 4,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := candy(tt.args.ratings); got != tt.want {
				t.Errorf("candy() = %v, want %v", got, tt.want)
			}
		})
	}
}
