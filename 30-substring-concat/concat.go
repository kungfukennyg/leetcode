package concat

func findSubstring(in string, words []string) []int {
	numWords := len(words)
	if numWords < 1 || len(in) < 1 || len(in) < len(words[0]) {
		return []int{}
	}

	wordCounts := make(map[string]int16, numWords)
	for _, w := range words {
		wordCounts[w]++
	}

	seen := make(map[string]int16, numWords)
	queue := make([]string, 0, numWords)
	var curIdx, queueLen int
	// pre-set capacity to avoid inner-loop allocations
	matching := make([]int, 0, 256)
	window := len(words[0])
	totalMatching := 0
	for i := 0; i < window; i++ {
		if i > 0 {
			curIdx, queueLen = 0, 0
			queue = queue[:0]
			clear(seen)
		}
		for j := i; j+window <= len(in); j += window {
			sub := in[j : j+window]
			count, matches := wordCounts[sub]
			if !matches {
				curIdx, queueLen = 0, 0
				queue = queue[:0]
				clear(seen)
				continue
			}

			for seen[sub] == count {
				q := queue[curIdx]
				seen[q]--
				curIdx++
				queueLen--
			}
			seen[sub]++
			queue = append(queue, sub)
			queueLen++
			if queueLen == numWords {
				matching = append(matching, j-(numWords-1)*window)
				totalMatching++
			}
		}
	}

	return matching[:totalMatching]
}
