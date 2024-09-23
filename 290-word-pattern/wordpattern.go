package wordpattern

import "strings"

func wordPattern(pattern string, str string) bool {
	if len(pattern) == 0 || len(str) == 0 {
		return false
	}

	words := strings.Split(str, " ")
	if len(words) != len(pattern) {
		return false
	}

	patternToStr := map[rune]string{}
	strToPattern := map[string]rune{}
	for i, s := range words {
		p := rune(pattern[i])
		if ss, ok := patternToStr[p]; ok {
			if s != ss {
				return false
			}
		}
		if pp, ok := strToPattern[s]; ok {
			if pp != p {
				return false
			}
		}
		patternToStr[p] = s
		strToPattern[s] = p
	}

	return true
}

func wordPattern2(pattern string, str string) bool {
	patternMap := map[rune]int{}
	strMap := map[string]int{}

	for _, p := range pattern {
		patternMap[p] += 1
	}

	for _, word := range strings.Split(str, " ") {
		count, ok := strMap[word]
		if !ok {
			count = 0
		}
		strMap[word] = count + 1
	}

OUTER:
	for c, pCount := range patternMap {
		for s, sCount := range strMap {
			if pCount == sCount {
				delete(patternMap, c)
				delete(strMap, s)
				continue OUTER
			}
		}

		return false
	}

	return true
}
