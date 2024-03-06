package majority

// majorityElement implements the Boyer Moore Majority Voting Algorithm to find the
// majority element that has more than N/2 occurences where N = len(input)
// see: https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore_majority_vote_algorithm
func majorityElement(nums []int) int {
	votes, candidate := 0, -1
	for _, n := range nums {
		if votes == 0 {
			candidate = n
			votes = 1
		} else {
			if n == candidate {
				votes++
			} else {
				votes--
			}
		}
	}

	return candidate
}
