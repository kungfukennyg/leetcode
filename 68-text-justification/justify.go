package justify

import (
	"strings"
)

type line struct {
	words    []string
	len      int
	maxWidth int
}

func (l *line) fits(word string) bool {
	// length + 1 space per word + new word
	return l.len+len(l.words)+len(word) <= l.maxWidth
}

func (l *line) add(word string) bool {
	l.words = append(l.words, word)
	l.len += len(word)
	return l.len == l.maxWidth
}

func (l *line) spaces() int {
	return l.maxWidth - l.len
}

func (l *line) justify(lastLine bool) string {
	spaces := l.spaces()
	numWords := len(l.words)
	if numWords == 1 || lastLine {
		return l.leftJustify(spaces)
	}

	var sb strings.Builder
	sb.Grow(l.maxWidth)
	for i, word := range l.words {
		sb.WriteString(word)
		var distribute int
		if i+1 < numWords {
			distribute = spaces / (numWords - 1)
		} else {
			distribute = 0
		}
		sb.WriteString(spaceBufs[distribute])
		remainder := spaces % (numWords - 1)
		if remainder > 0 {
			sb.WriteString(spaceBufs[1])
			spaces--
		}
	}
	return sb.String()
}

func (l *line) leftJustify(spaces int) string {
	var sb strings.Builder
	sb.Grow(l.maxWidth)
	singleSpace := spaceBufs[1]
	for i := 0; i < len(l.words)-1 && spaces > 0; i++ {
		sb.WriteString(l.words[i])
		sb.WriteString(singleSpace)
		spaces--
	}
	sb.WriteString(l.words[len(l.words)-1])
	if spaces > 0 {
		sb.WriteString(spaceBufs[spaces])
	}
	return sb.String()
}

// index = num of spaces
var spaceBufs []string

func fullJustify(words []string, maxWidth int) []string {
	// pre-allocate strings of spaces to avoid looping concatenation
	spaceBufs = make([]string, maxWidth+1)
	var sb strings.Builder
	spaceBufs[0] = ""
	// constraint: min word length = 1, max = maxWidth - 1
	for i := 1; i < maxWidth; i++ {
		sb.WriteRune(' ')
		spaceBufs[i] = sb.String()
	}

	out := []string{}
	cur := line{words: []string{}, maxWidth: maxWidth}
	for i := 0; i < len(words); i++ {
		word := words[i]
		if fits := cur.fits(word); fits && cur.add(word) {
			// added word to current line and line is full
			out = append(out, cur.justify(i == len(words)))
			cur = line{words: []string{}, maxWidth: maxWidth}
		} else if !fits {
			// word doesn't fit, next line
			out = append(out, cur.justify(i == len(words)))
			cur = line{words: []string{}, maxWidth: maxWidth}
			cur.add(word)
		}
		if i+1 == len(words) && cur.len > 0 {
			// hit the end of our words, justify last line
			out = append(out, cur.justify(i+1 == len(words)))
		}
	}
	return out
}

// 83%/1ms, 34%/2.51MB
func fullJustifyTwoPass(words []string, maxWidth int) []string {
	lines := []line{}
	// first split the input into words we can fit on each line
	cur := line{words: []string{}, maxWidth: maxWidth}
	for i := 0; i < len(words); i++ {
		word := words[i]
		if fits := cur.fits(word); fits && cur.add(word) {
			lines = append(lines, cur)
			cur = line{words: []string{}, maxWidth: maxWidth}
		} else if !fits {
			lines = append(lines, cur)
			cur = line{words: []string{}, maxWidth: maxWidth}
			cur.add(word)
		}
		if i+1 == len(words) && cur.len > 0 {
			lines = append(lines, cur)
		}
	}

	// now calculate the spaces needed per line
	out := make([]string, len(lines))
	for j, line := range lines {
		spaces := line.spaces()

		if len(line.words) == 1 || j+1 == len(lines) {
			out[j] = leftJustify(line.words, spaces)
			continue
		}

		out[j] = fullyJustify(line.words, spaces)
	}
	return out
}

func fullyJustify(words []string, spaces int) string {
	// add spacing
	for spaces > 0 {
		for i := 0; i < len(words)-1 && spaces > 0; i++ {
			words[i] += " "
			spaces--
		}
	}

	str := ""
	for i := 0; i < len(words); i++ {
		str += words[i]
	}
	return str
}

func leftJustify(words []string, spaces int) string {
	str := ""
	for i := 0; i < len(words)-1 && spaces > 0; i++ {
		str += words[i] + " "
		spaces--
	}
	str += words[len(words)-1]
	for spaces != 0 {
		str += " "
		spaces--
	}
	return str
}

// 83%/1ms, 34%/2.5MB
func fullJustifyThreePass(words []string, maxWidth int) []string {
	lines := make([][]string, 1)
	line := 0
	// first split the input into words we can fit on each line
	cur := []string{}
	lines[0] = cur
	for i := 0; i < len(words); i++ {
		// first word will always fit
		if len(cur) == 0 {
			cur = append(cur, words[i])
			continue
		}

		word := words[i]
		lineLen := lineLen(cur)
		// at least 1 space per word
		lineLen += len(cur)
		lineLen += len(word)
		if lineLen <= maxWidth {
			cur = append(cur, word)
			if lineLen == maxWidth {
				lines[line] = cur
				line++
				cur = []string{}
				lines = append(lines, cur)
			}
		} else {
			lines[line] = cur
			line++
			cur = []string{word}
			lines = append(lines, cur)
		}
	}
	if len(cur) > 0 {
		lines[line] = cur
	}
	if len(lines[line]) == 0 {
		lines = lines[:len(lines)-1]
	}

	// now calculate the spaces needed per line
	out := make([]string, len(lines))
	for j, words := range lines {
		wordsLen := lineLen(words)
		spaces := maxWidth - wordsLen
		if len(words) == 1 {
			// if only one word, all space goes to the right
			out[j] += words[0] + createSpaces(spaces)
			continue
		}

		// last line is left justified
		if j+1 == len(lines) {
			out[j] = leftJustify(words, spaces)
			continue
		}

		out[j] = fullyJustify(words, spaces)
	}
	return out
}

func createSpaces(n int) string {
	s := ""
	for range n {
		s += " "
	}
	return s
}

func lineLen(words []string) int {
	c := 0
	for _, w := range words {
		c += len(w)
	}
	return c
}
