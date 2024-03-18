package sum

const target = 0

func threeSum(nums []int) [][]int {
	// sort in increasing order to make processing easier
	sort(nums)

	triplets := [][]int{}
	for i, n1 := range nums {
		// we can skip any further elements since the slice
		// is sorted
		if n1 > target {
			break
		}
		// skip repeats
		if i > 0 && n1 == nums[i-1] {
			continue
		}

		for j, k := i+1, len(nums)-1; j < k; {
			n2, n3 := nums[j], nums[k]
			sum := n1 + n2 + n3
			if sum < target {
				j++
				continue
			} else if sum > target {
				k--
				continue
			}

			triplets = append(triplets, []int{n1, n2, n3})
			for j < k && nums[j] == n2 {
				j++
			}
			for j < k && nums[k] == n3 {
				k--
			}
		}
	}

	return triplets
}

// -10^5 <= nums[i] <= 10^5
const minValue int32 = -100_000
const mask = 0xFFFF_FFFF

func compact(i, j, k int32) int64 {
	ii := int64(i - minValue)
	jj := int64(j - minValue)
	kk := int64(k - minValue)
	// each value will fit into about 2^18 bytes,
	// shift a few extra to ensure there's no overlap
	return (ii << 42) | (jj << 21) | kk&mask
}

func unpack(m int64) (int32, int32, int32) {
	i := int32(((m >> 42) & 0x1FFF_FF)) + minValue
	j := int32(((m >> 21) & 0x1FFF_FF)) + minValue
	k := int32((m & 0x1FFF_FF)) + minValue
	return i, j, k
}

func sortTriplets(i, j, k int) (int32, int32, int32) {
	maxIdx := max(i, j)
	maxIdx = max(maxIdx, k)
	minIdx := min(i, j)
	minIdx = min(minIdx, k)
	var mid int
	if (maxIdx == i && minIdx == j) || (maxIdx == j && minIdx == i) {
		mid = k
	} else if (maxIdx == k && minIdx == j) || (minIdx == k && maxIdx == j) {
		mid = i
	} else {
		mid = j
	}

	return int32(minIdx), int32(mid), int32(maxIdx)
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

func threeSumMap(nums []int) [][]int {
	// sort in increasing order to make processing easier
	sort(nums)

	unique := map[int64]struct{}{}
	for i, n1 := range nums {
		// we can skip any further elements since the slice
		// is sorted; if n1 is already greater than our target,
		// nothing past this will ever sum to target.
		if n1 > target {
			break
		}
		// skip repeats
		if i > 0 && n1 == nums[i-1] {
			continue
		}

		for j, k := i+1, len(nums)-1; j < k; {
			n2, n3 := nums[j], nums[k]
			sum := n1 + n2 + n3
			if sum < target {
				j++
				continue
			} else if sum > target {
				k--
				continue
			}

			unique[compact(sortTriplets(n1, n2, n3))] = struct{}{}
			for j < k && nums[j] == n2 {
				j++
			}
			for j < k && nums[k] == n3 {
				k--
			}
		}
	}

	triplets := [][]int{}
	for m := range unique {
		a, b, c := unpack(m)
		triplets = append(triplets, []int{int(a), int(b), int(c)})
	}
	return triplets
}
