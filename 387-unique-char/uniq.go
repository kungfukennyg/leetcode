package uniq

import "math"

func firstUniqChar(in string) int {
	// A 97 Z 122
	chars := make([]int, 26)
	for i, s := range in {
		idx := s - 'a'
		if chars[idx] == 0 {
			chars[idx] = -(i + 1)
		} else {
			chars[idx] = i + 1
		}
	}
	lowest := math.MaxInt
	for _, s := range chars {
		if s < 0 {
			lowest = min(lowest, int(-s))
		}
	}
	if lowest == math.MaxInt {
		return -1
	} else {
		return lowest - 1
	}
}
