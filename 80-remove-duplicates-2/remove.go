package removeduplicates2

func removeDuplicates(nums []int) int {
	counts := map[int]int{}
	j := 0
	for i := 0; i < len(nums); i++ {
		a := nums[i]
		if c, ok := counts[a]; !ok || c+1 <= 2 {
			nums[j] = nums[i]
			j++
			if !ok {
				counts[a] = 1
				continue
			}
		}

		counts[a] = counts[a] + 1
	}

	nums = nums[:j]
	for i := 0; i < cap(nums)-j; i++ {
		nums = append(nums, -1)
	}
	return j
}
