package twosum2

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{numbers: []int{3, 24, 50, 79, 88, 150, 345}, target: 200},
			want: []int{3, 6},
		},
		{
			args: args{numbers: []int{5, 25, 75}, target: 100},
			want: []int{2, 3},
		},
		{
			args: args{numbers: []int{2, 7, 11, 15}, target: 9},
			want: []int{1, 2},
		},
		{
			args: args{numbers: []int{0, 0, 3, 4}, target: 0},
			want: []int{1, 2},
		},
		{
			args: args{numbers: []int{2, 3, 4}, target: 6},
			want: []int{1, 3},
		},
		{
			args: args{numbers: []int{-1, 0}, target: -1},
			want: []int{1, 2},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := twoSum(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
