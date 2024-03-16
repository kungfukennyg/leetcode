package palindrome

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		l, r := lower(rune(s[left])), lower(rune(s[right]))
		if l == ' ' {
			left++
			continue
		}
		if r == ' ' {
			right--
			continue
		}
		if l != r {
			return false
		}

		left++
		right--
	}
	return true
}

// lower returns the lowercase letter if r is a letter, or otherwise an empty
// space if r is non-alphanumeric
func lower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		r += 32
	}
	if (r < '0' || r > '9') && (r < 'a' || r > 'z') {
		return ' '
	}
	return r
}
