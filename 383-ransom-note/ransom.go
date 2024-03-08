package ransom

func canConstruct(ransom string, mag string) bool {
	if ransom == mag {
		return true
	}

	counts := map[rune]int{}
	for _, l := range mag {
		counts[l] += 1
	}

	ransomCounts := map[rune]int{}
	for _, l := range ransom {
		ransomCounts[l] += 1
	}
	for c, v := range ransomCounts {
		m := counts[c]
		if v > m {
			return false
		}
	}
	return true
}
