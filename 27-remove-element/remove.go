package remove

func removeElement(nums []int, val int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}

	nums = nums[:j]
	counted := len(nums)
	for i := 0; i < j; i++ {
		nums = append(nums, -1)
	}

	return counted
}
