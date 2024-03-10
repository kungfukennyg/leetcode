package array

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{[]int{1, 2, 3, 4}},
			want: []int{24, 12, 8, 6},
		},
		{
			args: args{[]int{-1, 1, 0, -3, 3}},
			want: []int{0, 0, 9, 0, 0},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := productExceptSelf(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_productExceptSelf(b *testing.B) {
	inputs := []int{10, 100, 1000, 10_000}
	for i, in := range inputs {
		input := make([]int, in)
		for i := range input {
			input[i] = -30 + rand.Intn(31)
		}

		b.Run(fmt.Sprint(i), func(b *testing.B) {
			for range b.N {
				productExceptSelf(input)
			}
		})
	}
}
