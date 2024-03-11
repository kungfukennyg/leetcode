# Intuition
My initial approach to this problem was a brute-force nested loop solution that tried to look ahead to determine the end of a range between two peaks. This worked for some test cases but required some kludges and became really long and cumbersome to reason about as I added more tests. I next tried pre-calculating the ranges water could fit between with a a single loop but skipping backwards when a range was incomplete by the end, and ran into similar issues with some test cases breaking the logic. 

# Approach
I finally settled on finding the highest point of the height map and using that to perform two loops, calculating the collected water from the left and right of the highest point. I also tested doing this in one loop, checking the right and left of the highest point simultaneously. This proved similar, albeit slower, on low-N inputs but saw a significant slowdown on high-N; see benchmark results below.

# Complexity
- Time complexity:
O(3n)

- Space complexity:
O(1)

# Benchmarks
```
goos: linux
goarch: amd64
pkg: example.com/leetcode/42-trapping-rain-water
cpu: 13th Gen Intel(R) Core(TM) i9-13900K
Benchmark_trap/0_10-32         	100000000	        11.16 ns/op	       0 B/op	       0 allocs/op
Benchmark_trap/1_100-32        	16217552	        73.68 ns/op	       0 B/op	       0 allocs/op
Benchmark_trap/2_1000-32       	 1477731	       811.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_trap/3_10000-32      	  165996	      7165 ns/op	       0 B/op	       0 allocs/op
Benchmark_trap/4_100000-32     	   15754	     74808 ns/op	       0 B/op	       0 allocs/op
Benchmark_trap/5_1000000-32    	    1556	    780880 ns/op	       0 B/op	       0 allocs/op
Benchmark_trapSingleLoop/0_10-32         	95475578	        12.93 ns/op	       0 B/op	       0 allocs/op
Benchmark_trapSingleLoop/1_100-32        	13108610	        90.99 ns/op	       0 B/op	       0 allocs/op
Benchmark_trapSingleLoop/2_1000-32       	 1299896	       919.8 ns/op	       0 B/op	       0 allocs/op
Benchmark_trapSingleLoop/3_10000-32      	  132109	      9043 ns/op	       0 B/op	       0 allocs/op
Benchmark_trapSingleLoop/4_100000-32     	   13566	     87908 ns/op	       0 B/op	       0 allocs/op
Benchmark_trapSingleLoop/5_1000000-32    	    1299	    911320 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	example.com/leetcode/42-trapping-rain-water	18.284s
```

# Code
```
func trap(height []int) int {
	// find the highest point. it's fine if there are multiple highest points
	// and we flag the last occurence
	highest, highestIdx := -1, -1
	for i, h := range height {
		if h > highest {
			highest = h
			highestIdx = i
		}
	}

	// process to the left of the highest point
	water := 0
	highest = -1
	for i := 0; i < highestIdx; i++ {
		h := height[i]
		if h < highest {
			water += highest - h
		} else if h >= highest {
			highest = h
		}
	}

	// ditto for the right
	highest = -1
	for i := len(height) - 1; i > highestIdx; i-- {
		h := height[i]
		if h < highest {
			water += highest - h
		} else if h >= highest {
			highest = h
		}
	}
	return water
}

func trapSingleLoop(height []int) int {
	// find the highest point. it's fine if there are multiple highest points
	// and we flag the last occurence
	highest, highestIdx := -1, -1
	for i, h := range height {
		if h > highest {
			highest = h
			highestIdx = i
		}
	}

	water := 0
	leftHigh, rightHigh := -1, -1
	for i := 0; i < len(height); i++ {
		add := 0
		// left
		if i < highestIdx {
			h := height[i]
			if h < leftHigh {
				water += leftHigh - h
			} else if h >= leftHigh {
				leftHigh = h
			}
		}

		// right
		if len(height)-1-i > highestIdx {
			h := height[len(height)-1-i]
			if h < rightHigh {
				water += rightHigh - h
			} else if h >= rightHigh {
				rightHigh = h
			}
		}

		water += add
	}
	return water
}
```

# Tests
```
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

```