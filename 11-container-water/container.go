package container

func maxArea(heights []int) int {
	maxArea := -1
	l, r := 0, len(heights)-1
	for l != r {
		width := r - l
		height := min(heights[l], heights[r])
		area := width * height
		maxArea = max(maxArea, area)
		if heights[l] <= heights[r] {
			l++
		} else if heights[l] > heights[r] {
			r--
		}
	}

	return maxArea
}
