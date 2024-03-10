package candy

// two pass O(2n)
func candy(ratings []int) int {
	give := make([]int, len(ratings))
	total := len(ratings)

	process := func(i, next int, ratings []int, give []int) int {
		added := 0
		c, rc := give[i], give[next]
		if c == 0 {
			c = 1
		}
		if rc == 0 {
			rc = 1
		}
		if ratings[i] > ratings[next] && c <= rc {
			add := rc + 1
			added += add - c
			give[i] = add
		}
		return added
	}

	// neighbors left to right
	for i := 1; i < len(ratings); i++ {
		total += process(i, i-1, ratings, give)
	}

	// same process, right to left
	for i := len(ratings) - 2; i >= 0; i-- {
		total += process(i, i+1, ratings, give)
	}
	return total
}
