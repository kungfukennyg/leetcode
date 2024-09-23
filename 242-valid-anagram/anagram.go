package anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sCount := map[rune]int{}
	for _, c := range s {
		sCount[c] += 1
	}

	for _, c := range t {
		count, ok := sCount[c]
		if !ok {
			return false
		}

		count--
		if count < 0 {
			return false
		}
		sCount[c] = count
	}

	return true
}
