package array

func productExceptSelf(nums []int) []int {
	left, right := 1, 1
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = 1
	}

	for i := 0; i < n; i++ {
		ans[i] *= left
		left = left * nums[i]
		ans[n-i-1] *= right
		right = right * nums[n-i-1]
	}

	return ans
}
