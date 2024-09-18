package matrixzero

import "math"

const rewriteFlag = math.MaxInt64

func setZeroes(matrix [][]int) {
	for y, row := range matrix {
		rewriteRow := false
		for x, elem := range row {
			// found a zero
			if elem == 0 {
				rewriteRow = true
				// rewrite this column now
				for y := 0; y < len(matrix); y++ {
					// if it's already a zero, skip so we don't miss it later
					if matrix[y][x] != 0 {
						matrix[y][x] = rewriteFlag
					}
				}
			}
		}
		// now rewrite the entire row
		if rewriteRow {
			for x := 0; x < len(matrix[0]); x++ {
				matrix[y][x] = rewriteFlag
			}
		}
	}

	// all element's marked, now write the zeroes
	for y, row := range matrix {
		for x, elem := range row {
			if elem == rewriteFlag {
				matrix[y][x] = 0
			}
		}
	}
}
