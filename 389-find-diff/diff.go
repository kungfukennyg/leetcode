package diff

func findTheDifference(s, t string) byte {
	a, b := []byte(s), []byte(t)
OUTER:
	for len(b) > 0 {
		s := b[0]
		for i, v := range a {
			if s == byte(v) {
				last := a[len(a)-1]
				a[i] = last
				a = a[:len(a)-1]
				b = b[1:]
				continue OUTER
			}
		}
		return s
	}

	return 0
}
