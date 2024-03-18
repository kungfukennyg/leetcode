package substring

func lengthOfLongestSubstring(in string) int {
	if len(in) == 1 {
		return 1
	}

	longest := 0
OUTER:
	for i, c := range in {
		// ascii 32-127 = 95
		chars := make([]int, 95)
		chars[c-32] = 1
		j := i + 1
		for ; j < len(in); j++ {
			idx := in[j] - 32
			if chars[idx] > 0 {
				longest = max(longest, j-i)
				continue OUTER
			}
			chars[idx]++
		}

		longest = max(longest, j-i)
	}
	return longest
}
