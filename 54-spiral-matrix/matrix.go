package matrix

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	width, height := len(matrix[0]), len(matrix)
	elems := width * height
	ret := make([]int, 0)

	x, y := 0, 0
	sp := newSpiral(DirRight, width, height)
	for len(ret) < elems {
		elem := matrix[y][x]
		ret = append(ret, elem)
		x, y = sp.next(x, y)
	}
	return ret
}

type spiral struct {
	dir
	minWidth  int
	minHeight int
	maxWidth  int
	maxHeight int
}

func newSpiral(startDir dir, width, height int) *spiral {
	return &spiral{
		dir:       startDir,
		maxWidth:  width,
		maxHeight: height,
	}
}

func (s *spiral) next(x, y int) (int, int) {
	if s.dir == DirRight && x+1 >= s.maxWidth {
		s.minHeight++
		s.dir = s.dir.next()
	} else if s.dir == DirDown && y+1 >= s.maxHeight {
		s.maxWidth--
		s.dir = s.dir.next()
	} else if s.dir == DirLeft && x-1 < s.minWidth {
		s.maxHeight--
		s.dir = s.dir.next()
	} else if s.dir == DirUp && y-1 < s.minHeight {
		s.minWidth++
		s.dir = s.dir.next()
	}

	return s.dir.move(x, y)
}

type dir string

const (
	DirRight dir = "right"
	DirLeft  dir = "left"
	DirDown  dir = "down"
	DirUp    dir = "up"
)

func (d dir) next() dir {
	switch d {
	case DirDown:
		return DirLeft
	case DirLeft:
		return DirUp
	case DirRight:
		return DirDown
	case DirUp:
		return DirRight
	default:
		panic(fmt.Sprintf("unexpected matrix.dir: %#v", d))
	}
}

func (d dir) move(x, y int) (int, int) {
	switch d {
	case DirDown:
		return x, y + 1
	case DirLeft:
		return x - 1, y
	case DirRight:
		return x + 1, y
	case DirUp:
		return x, y - 1
	default:
		panic(fmt.Sprintf("unexpected matrix.dir: %#v", d))
	}
}
