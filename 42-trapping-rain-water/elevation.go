package elevation

import (
	"fmt"
	"strings"
)

func trap(height []int) int {
	// find the highest point. it's fine if there are multiple highest points
	// and we flag the last occurence
	highest, highestIdx := -1, -1
	for i, h := range height {
		if h > highest {
			highest = h
			highestIdx = i
		}
	}

	// process to the left of the highest point
	water := 0
	highest = -1
	for i := 0; i < highestIdx; i++ {
		h := height[i]
		if h < highest {
			water += highest - h
		} else if h >= highest {
			highest = h
		}
	}

	// ditto for the right
	highest = -1
	for i := len(height) - 1; i > highestIdx; i-- {
		h := height[i]
		if h < highest {
			water += highest - h
		} else if h >= highest {
			highest = h
		}
	}
	return water
}

func trapSingleLoop(height []int) int {
	// find the highest point. it's fine if there are multiple highest points
	// and we flag the last occurence
	highest, highestIdx := -1, -1
	for i, h := range height {
		if h > highest {
			highest = h
			highestIdx = i
		}
	}

	water := 0
	leftHigh, rightHigh := -1, -1
	for i := 0; i < len(height); i++ {
		add := 0
		// left
		if i < highestIdx {
			h := height[i]
			if h < leftHigh {
				water += leftHigh - h
			} else if h >= leftHigh {
				leftHigh = h
			}
		}

		// right
		if len(height)-1-i > highestIdx {
			h := height[len(height)-1-i]
			if h < rightHigh {
				water += rightHigh - h
			} else if h >= rightHigh {
				rightHigh = h
			}
		}

		water += add
	}
	return water
}

func toString(heights []int) string {
	maxHeight := -1
	for _, h := range heights {
		maxHeight = max(maxHeight, h)
	}
	if maxHeight == -1 {
		return ""
	}

	var sb strings.Builder
	for i := maxHeight; i > 0; i-- {
		sb.WriteString(fmt.Sprintf("%d ", i))

		for _, h := range heights {
			if h >= i {
				sb.WriteString("[]  ")
			} else {
				sb.WriteString("    ")
			}
		}
		sb.WriteString("\n")
	}
	for i := range heights {
		sb.WriteString(fmt.Sprintf("  %d ", i))
	}
	return sb.String()
}
