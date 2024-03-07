package hindex

func hIndex(citations []int) int {
	sort(citations)

	h := 0
	for i := len(citations) - 1; i >= 0; i-- {
		// for each publication, keep a count of publications cited
		// up until n < times cited
		n := citations[i]
		if n > h {
			h++
		} else {
			break
		}
	}

	return h
}

func sort(s []int) {
	// bubble sort in non-decreasing order
	var tmp int
	for i := 0; i < len(s)-1; i++ {
		swapped := false
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				// swap
				tmp = s[j]
				s[j] = s[j+1]
				s[j+1] = tmp
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
