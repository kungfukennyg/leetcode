package shuffle

import "math/rand"

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

func (s *Solution) Reset() []int {
	return s.nums
}

func (s *Solution) Shuffle() []int {
	len := len(s.nums)
	ret := make([]int, len)
	copy(ret, s.nums)
	for i := 0; i < len; i++ {
		j := i + rand.Intn(len-i)
		ret[i], ret[j] = ret[j], ret[i]
	}
	return ret
}
