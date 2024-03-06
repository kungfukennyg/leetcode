package removedupes

func removeDuplicates(nums []int) int {
	elements := map[int]struct{}{}
	j := 0
	for i := 0; i < len(nums); i++ {
		a := nums[i]
		// if element is not a duplicate then update our slice in place
		if _, ok := elements[a]; !ok {
			nums[j] = nums[i]
			j++
			elements[a] = struct{}{}
		}
	}
	nums = nums[:j]
	for i := 0; i < j; i++ {
		nums = append(nums, -1)
	}
	return len(elements)
}
