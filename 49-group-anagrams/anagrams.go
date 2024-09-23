package groupanagrams

import (
	"slices"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	ret := map[string]struct{}{}
	for _, word := range strs {
		if len(word) == 0 {
			continue
		}
		chars := map[rune]int{}
		for _, s := range word {
			chars[s] += 1
		}
		anagrams := []string{word}
		for _, str := range strs {
			if word == str {
				continue
			}
			if isAnagram(str, chars) {
				anagrams = append(anagrams, str)
			}
		}
		slices.Sort(anagrams)
		line := strings.Join(anagrams, " ")
		ret[line] = struct{}{}
	}

	out := [][]string{}
	for r := range ret {
		out = append(out, strings.Split(r, " "))
	}
	for _, w := range strs {
		if len(w) == 0 {
			out = append(out, []string{""})
		}
	}
	return out
}

func isAnagram(str string, chars map[rune]int) bool {
	strChars := map[rune]int{}
	for _, c := range str {
		if _, ok := chars[c]; !ok {
			return false
		}

		strChars[c] += 1
		if strChars[c] > chars[c] {
			return false
		}
	}

	for c, count := range strChars {
		if count != chars[c] {
			return false
		}
	}

	return true
}
