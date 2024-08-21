package sudoku

const BLANK byte = '.'
const SIZE = 9

type set map[byte]struct{}

func isValidSudoku(rows [][]byte) bool {
	rowSeen := set{}
	colSeen := map[int]set{}
	subSeen := map[int]set{}

	for y, row := range rows {
		clear(rowSeen)

		for x, cell := range row {
			// first row, setup column
			if y == 0 {
				colSeen[x] = set{}
			}

			// calculate 3x3 subchunk position
			subX := x / 3
			subY := y / 3
			subChunk := subX + (subY * 3)
			if _, ok := subSeen[subChunk]; !ok {
				subSeen[subChunk] = set{}
			}

			// we don't care if a cell is blank, we're only validating
			// filled cells
			if cell == BLANK {
				continue
			}

			// check our sets for duplicates

			// column
			if _, ok := colSeen[x][cell]; ok {
				return false
			}

			// row
			if _, ok := rowSeen[cell]; ok {
				return false
			}

			// 3x3
			if _, ok := subSeen[subChunk][cell]; ok {
				return false
			}

			colSeen[x][cell] = struct{}{}
			rowSeen[cell] = struct{}{}
			subSeen[subChunk][cell] = struct{}{}
		}
	}

	return true
}
