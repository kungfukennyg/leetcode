package container

import (
	"fmt"
	"strings"
)

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

func toString(height []int) string {
	var sb strings.Builder
	highest := -1
	for _, h := range height {
		highest = max(highest, h)
	}

	for i := highest; i > 0; i-- {
		for _, h := range height {
			if h >= i {
				sb.WriteString("x ")
			} else {
				sb.WriteString("  ")
			}
		}
		sb.WriteString("\n")
	}
	for i := range height {
		sb.WriteString(fmt.Sprint(i) + " ")
	}
	return sb.String()
}

func maxAreaBroken(height []int) int {
	highest := -1
	highestIdx := -1
	for i, h := range height {
		if h > highest {
			highest = h
			highestIdx = i
		}
	}
	maxArea := -1
	// check left
	for i := highestIdx - 1; i >= 0; i-- {
		h := height[i]
		width := highestIdx - i
		height := min(h, highest)
		area := width * height
		maxArea = max(area, width*height)
	}
	// check right
	for i := highestIdx + 1; i < len(height); i++ {
		h := height[i]
		width := i - highestIdx
		height := min(h, highest)
		area := width * height
		maxArea = max(area, width*height)
	}
	return maxArea
}
