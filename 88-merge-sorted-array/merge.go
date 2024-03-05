package merge

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) != m+n {
		return
	}
	if len(nums2) != n {
		return
	}

	if len(nums2) == 0 {
		return
	}

	// insert all elements of 2 into 1 in zero slots
	var j int
	for i, a := range nums1 {
		if a != 0 {
			continue
		}

		nums1[i] = nums2[j]
		j++
		if j >= n {
			break
		}
	}

	// bubble sort in non-decreasing order
	var tmp int
	for i := 0; i < len(nums1)-1; i++ {
		swapped := false
		for j := 0; j < len(nums1)-i-1; j++ {
			if nums1[j] > nums1[j+1] {
				// swap
				tmp = nums1[j]
				nums1[j] = nums1[j+1]
				nums1[j+1] = tmp
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
