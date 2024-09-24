package sequence

import (
	"math"
	"slices"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slices.Sort(nums)

	highest := math.MinInt64
	consecutive := 1
	prev := math.MinInt64
	for _, n := range nums {
		if prev == math.MinInt64 {
			prev = n
			continue
		}

		if n == prev {
			continue
		} else if n == prev+1 {
			consecutive++
		} else {
			highest = max(highest, consecutive)
			consecutive = 1
		}

		prev = n
	}

	return max(highest, consecutive)
}
