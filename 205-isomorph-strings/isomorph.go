package isomorph

func isIsomorphic(str string, pattern string) bool {
	if len(str) != len(pattern) {
		return false
	}
	if len(str) == 0 {
		return true
	}

	strMap := map[rune]rune{}
	patternMap := map[rune]rune{}
	for i, c := range str {
		p := rune(pattern[i])
		if pp, ok := strMap[c]; ok && p != pp {
			return false
		}
		strMap[c] = p

		if cc, ok := patternMap[p]; ok && c != cc {
			return false
		}

		patternMap[p] = c
	}

	return true
}
