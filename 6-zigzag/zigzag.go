package zigzag

// ~89%/2ms, 67%/6.12MB
func convert(in string, rows int) string {
	if rows <= 1 {
		return in
	}

	lines := make([][]byte, rows)
	y := 0
	edge := false
	for i := 0; i < len(in); i++ {
		if y == 0 || y == rows-1 {
			edge = !edge
		}
		lines[y] = append(lines[y], in[i])
		if edge {
			y++
		} else {
			y--
		}
	}
	out := make([]byte, len(in))
	i := 0
	for _, line := range lines {
		for _, b := range line {
			out[i] = b
			i++
		}
	}
	return string(out)
}
