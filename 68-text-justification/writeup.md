# Intuition
My initial approach to this solution involved three passes; the first to delineate words per line, the second to calculate spaces needed and concatenate into one string, and the final to write each string to the output buffer. This worked but the code was messy, and I theorized there was a lot of room for cleanup and optimization.

# Approach
I iterated on my initial solution by introducing a simple `line` type to denote a given line of text, with some internal state to track the length of the text within the line, and collapsed my final two passes into one, justifying the text and writing to the final output buffer as each line was completed. I then iterated on this approach again by collapsing into one single pass, writing internal justify and leftJustify functions for `line`, and using `strings.Builder` and a pre-allocated set of space strings ("", " ", "  ", etc..) to avoid extra iteration while inserting spaces into each line. This final implementation proved to be the fastest and take the least allocations, as demonstrated in the below benchmarks.

# Benchmarks
```
Running tool: /usr/local/go/bin/go test -benchmem -run=^$ -bench ^BenchmarkFullJustify$ example.com/leetcode/68-text-justification

goos: linux
goarch: amd64
pkg: example.com/leetcode/68-text-justification
cpu: 13th Gen Intel(R) Core(TM) i9-13900K
BenchmarkFullJustify/-15-32         	 1572250	       759.2 ns/op	    1128 B/op	      28 allocs/op
BenchmarkFullJustify/-30-32         	 1714070	       703.1 ns/op	    1368 B/op	      21 allocs/op
BenchmarkFullJustify/-50-32         	 1633910	       732.8 ns/op	    1928 B/op	      18 allocs/op
BenchmarkFullJustify/-100-32        	 1409625	       859.3 ns/op	    3176 B/op	      14 allocs/op
BenchmarkFullJustify/TwoPass-15-32  	  945253	      1222 ns/op	    1616 B/op	      60 allocs/op
BenchmarkFullJustify/TwoPass-30-32  	  993032	      1154 ns/op	    1672 B/op	      57 allocs/op
BenchmarkFullJustify/TwoPass-50-32  	  892004	      1328 ns/op	    2480 B/op	      62 allocs/op
BenchmarkFullJustify/TwoPass-100-32 	  946011	      1266 ns/op	    4104 B/op	      48 allocs/op
BenchmarkFullJustify/ThreePass-15-32         	  980575	      1192 ns/op	    1280 B/op	      59 allocs/op
BenchmarkFullJustify/ThreePass-30-32         	  982250	      1172 ns/op	    1528 B/op	      56 allocs/op
BenchmarkFullJustify/ThreePass-50-32         	  863595	      1366 ns/op	    2400 B/op	      61 allocs/op
BenchmarkFullJustify/ThreePass-100-32        	  876228	      1316 ns/op	    4056 B/op	      47 allocs/op
PASS
ok  	example.com/leetcode/68-text-justification	17.346s

func BenchmarkFullJustify(b *testing.B) {
	input := []string{"ask", "not", "what", "your", "country", "can", "do", "for", "you", "ask", "what", "you", "can", "do", "for", "your", "country"}
	widths := []int{15, 30, 50, 100}

	funcs := []struct {
		name string
		fn   func([]string, int) []string
	}{
		{
			name: "",
			fn:   fullJustify,
		},
		{
			name: "TwoPass",
			fn:   fullJustifyTwoPass,
		},
		{
			name: "ThreePass",
			fn:   fullJustifyThreePass,
		},
	}

	for _, tc := range funcs {
		for _, width := range widths {
			b.Run(fmt.Sprintf("%s-%d", tc.name, width), func(b *testing.B) {
				for range b.N {
					tc.fn(input, width)
				}
			})
		}
	}

```

# Complexity
- Time complexity:
O(n) overall where n is the number of words 

- Space complexity:
O(m) where m is the number of characters in the input text

# Code
```
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

```