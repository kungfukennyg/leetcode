package rotate

func rotate(nums []int, steps int) {
	if len(nums) == 1 {
		return
	}

	// calculate each elements new index based on current index + steps
	updated := make([]int, len(nums))
	for i, n := range nums {
		newIdx := i + steps
		for newIdx >= len(nums) {
			newIdx = newIdx - len(nums)
		}
		updated[newIdx] = n
	}

	copy(nums, updated)
}
