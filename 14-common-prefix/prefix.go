package prefix

import "math"

func longestCommonPrefix(strs []string) string {
	shortest := math.MaxInt
	for _, s := range strs {
		shortest = min(len(s), shortest)
	}
	if shortest == math.MaxInt || shortest == 0 {
		return ""
	}

	prefix := []rune{}
OUTER:
	for i := 0; i < shortest; i++ {
		for _, s := range strs {
			c := rune(s[i])
			if len(prefix) <= i {
				prefix = append(prefix, c)
			} else if prefix[i] != c {
				prefix = prefix[:len(prefix)-1]
				break OUTER
			}
		}
		if len(prefix) == i {
			break
		}
	}
	return string(prefix)
}
