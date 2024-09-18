package life

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_gameOfLife(t *testing.T) {
	type args struct {
		board [][]int
	}
	tests := []struct {
		args args
		want [][]int
	}{
		{
			args: args{[][]int{
				{0, 1, 0},
				{0, 0, 1},
				{1, 1, 1},
				{0, 0, 0},
			}},
			want: [][]int{
				{0, 0, 0},
				{1, 0, 1},
				{0, 1, 1},
				{0, 1, 0},
			},
		},
		{
			args: args{[][]int{
				{1, 1},
				{1, 0},
			}},
			want: [][]int{
				{1, 1},
				{1, 1},
			},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			gameOfLife(tt.args.board)
			if got := tt.args.board; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gameOfLife() = %v, want %v", got, tt.want)
			}
		})
	}
}
