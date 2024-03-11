package elevation

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func Test_trap(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{[]int{0, 1, 0, 3, 1, 0, 1, 3, 2, 1, 2, 1}},
			want: 9,
		},
		{
			args: args{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}},
			want: 6,
		},
		{
			args: args{[]int{4, 2, 3}},
			want: 1,
		},
		{
			args: args{[]int{4, 2, 0, 3, 2, 5}},
			want: 9,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			s := toString(tt.args.height)
			t.Logf("\n%s", s)
			if got := trap(tt.args.height); got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_trap(b *testing.B) {
	inputs := []int{10, 100, 1000, 10000, 100_000, 1_000_000}

	for i, in := range inputs {
		nums := make([]int, in)
		for n := 0; n < in; n++ {
			// constraints: 0 <= n < 10^5
			nums[n] = rand.Intn(int(math.Pow(10, 5)))
		}
		b.Run(fmt.Sprint(i, in), func(b *testing.B) {
			for range b.N {
				trap(nums)
			}
		})
	}
}
func Benchmark_trapSingleLoop(b *testing.B) {
	inputs := []int{10, 100, 1000, 10000, 100_000, 1_000_000}

	for i, in := range inputs {
		nums := make([]int, in)
		for n := 0; n < in; n++ {
			// constraints: 0 <= n < 10^5
			nums[n] = rand.Intn(int(math.Pow(10, 5)))
		}
		b.Run(fmt.Sprint(i, in), func(b *testing.B) {
			for range b.N {
				trapSingleLoop(nums)
			}
		})
	}
}
