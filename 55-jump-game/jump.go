package jump

func canJump(nums []int) bool {
	if len(nums) < 1 {
		return true
	}

	gas := 0
	for i, n := range nums {
		if gas < 0 {
			return false
		} else if n > gas {
			gas = n
		}
		if i+1 >= len(nums) {
			return true
		}
		gas -= 1
	}

	return false
}
