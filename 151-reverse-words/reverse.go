package reverse

func reverseWords(in string) string {
	// find leading/trailing spaces spaces
	bytes := []byte(in)
	start, end := 0, len(bytes)-1
	for start < len(bytes) && bytes[start] == ' ' {
		start++
	}
	for end > start && bytes[end] == ' ' {
		end--
	}
	bytes = bytes[start : end+1]

	// remove extra internal spaces in-place
	var prev byte
	j := 0
	for _, b := range bytes {
		if b != ' ' || (b == ' ' && prev != ' ') {
			bytes[j] = b
			j++
		}
		prev = b
	}
	bytes = bytes[:j]

	// reverse
	for i := len(bytes)/2 - 1; i >= 0; i-- {
		bytes[i], bytes[len(bytes)-i-1] = bytes[len(bytes)-i-1], bytes[i]
	}

	// now reverse each word's order in-place
	wordIdx := 0
	for i := 0; i <= len(bytes); i++ {
		if i == len(bytes) || bytes[i] == ' ' {
			for j := wordIdx; j < (wordIdx+i)/2; j++ {
				bytes[j], bytes[i-1-j+wordIdx] = bytes[i-1-j+wordIdx], bytes[j]
			}
			wordIdx = i + 1
		}
	}

	return string(bytes)
}
