package life

func gameOfLife(gameBoard [][]int) {
	if len(gameBoard) == 0 || len(gameBoard[0]) == 0 {
		return
	}

	board := board(gameBoard)
	board.step()
}

type board [][]int

func (b board) step() {
	neighbors := make([][]int, len(b))
	for y, row := range b {
		neighbors[y] = make([]int, len(row))
		for x := range row {
			neighbors[y][x] = b.neighbors(x, y)
		}
	}

	for y, row := range neighbors {
		for x, neighbor := range row {
			cell := b[y][x]
			if cell == 1 {
				if neighbor < 2 {
					b[y][x] = 0
				} else if neighbor > 3 {
					b[y][x] = 0
				}
			} else {
				if neighbor == 3 {
					b[y][x] = 1
				}
			}
		}
	}
}

func (b board) neighbors(x, y int) int {
	neighbors := 0
	height := len(b)
	width := len(b[0])
	// left
	if x > 0 && b[y][x-1] == 1 {
		neighbors++
	}
	// right
	if x+1 < width && b[y][x+1] == 1 {
		neighbors++
	}

	// above
	if y > 0 && b[y-1][x] == 1 {
		neighbors++
	}
	// below
	if y+1 < height && b[y+1][x] == 1 {
		neighbors++
	}

	// up diagonals
	if y > 0 {
		if x > 0 && b[y-1][x-1] == 1 {
			neighbors++
		}
		if x+1 < width && b[y-1][x+1] == 1 {
			neighbors++
		}
	}

	// down diagonals
	if y+1 < height {
		if x > 0 && b[y+1][x-1] == 1 {
			neighbors++
		}
		if x+1 < width && b[y+1][x+1] == 1 {
			neighbors++
		}
	}

	return neighbors
}
