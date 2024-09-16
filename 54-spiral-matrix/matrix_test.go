package matrix

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}},
			want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			args: args{[][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			}},
			want: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			args: args{[][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
				{21, 22, 23, 24, 25},
			}},
			want: []int{
				1, 2, 3, 4, 5,
				10, 15, 20, 25, 24,
				23, 22, 21, 16, 11,
				6, 7, 8, 9, 14,
				19, 18, 17, 12, 13},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := spiralOrder(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
