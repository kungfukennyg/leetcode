package subsequence

func isSubsequence(sub string, text string) bool {
	if len(sub) > len(text) {
		return false
	}
	if len(sub) == 0 {
		return true
	}

	j := 0
	for _, t := range text {
		if j == len(sub) {
			return true
		}
		if t == rune(sub[j]) {
			j++
			continue
		}
	}
	return j == len(sub)
}
