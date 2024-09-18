package matrixzero

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_setZeroes(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		args args
		want [][]int
	}{
		{
			args: args{[][]int{
				{0, 0, 0, 5},
				{4, 3, 1, 4},
				{0, 1, 1, 4},
				{1, 2, 1, 3},
				{0, 0, 1, 1},
			}},
			want: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 4},
				{0, 0, 0, 0},
				{0, 0, 0, 3},
				{0, 0, 0, 0},
			},
		},
		{
			args: args{[][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			}},
			want: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			args: args{[][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			}},
			want: [][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			setZeroes(tt.args.matrix)
			if got := tt.args.matrix; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
