package subarray

import "math"

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n > 0 && nums[0] >= target {
		return 1
	}

	left := make([]int, len(nums))
	left[0] = nums[0]
	lowest := math.MaxInt
	for i := 1; i < n; i++ {
		sum := nums[i] + left[i-1]
		left[i] = sum
		if sum < target {
			continue
		}

		j := 0
		for sum >= target {
			sum -= nums[j]
			j++
		}
		lowest = min(lowest, i+1-(j-1))
	}

	if lowest == math.MaxInt {
		lowest = 0
	}

	return lowest
}
