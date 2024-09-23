package twosum

func twoSum(nums []int, target int) []int {
	for i, n := range nums {
		for j, t := range nums {
			if i == j {
				continue
			}

			if n+t == target {
				return []int{i, j}
			}
		}
	}

	return nil
}
