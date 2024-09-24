package dupe2

type set map[int]struct{}

func containsNearbyDuplicate(nums []int, k int) bool {
	matches := map[int]set{}
	for i, a := range nums {
		if matchingIdxs, ok := matches[a]; ok {
			matchingIdxs[i] = struct{}{}
		} else {
			matches[a] = set{i: struct{}{}}
		}
	}

	for _, matchingIdxs := range matches {
		for i := range matchingIdxs {
			for j := range matchingIdxs {
				if i == j {
					continue
				}

				n := i - j
				if n < 0 {
					n *= -1
				}

				if n <= k {
					return true
				}
			}
		}
	}
	return false
}

func containsNearbyDuplicate_v2(nums []int, k int) bool {
	for i, a := range nums {
		for j, b := range nums {
			if i == j {
				continue
			}
			if j-i > k {
				continue
			}

			if a != b {
				continue
			}

			n := i - j
			if n < 0 {
				n *= -1
			}

			if n <= k {
				return true
			}
		}
	}
	return false
}

func containsNearbyDuplicate_v1(nums []int, k int) bool {
	for i, a := range nums {
		for j, b := range nums {
			if i == j {
				continue
			}

			if a != b {
				continue
			}

			n := i - j
			if n < 0 {
				n *= -1
			}

			if n <= k {
				return true
			}
		}
	}
	return false
}
