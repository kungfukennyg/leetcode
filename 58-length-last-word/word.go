package word

func lengthOfLastWord(in string) int {
	count := 0
	for i := len(in) - 1; i >= 0; i-- {
		s := in[i]
		if count > 0 && s == ' ' {
			break
		} else if s != ' ' {
			count++
		}
	}
	return count
}
