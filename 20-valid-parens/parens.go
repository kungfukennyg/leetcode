package parens

type bracket struct {
	pair rune
	open bool
}

var brackets map[rune]bracket = map[rune]bracket{
	'(': {
		pair: ')',
		open: true,
	},
	'{': {
		pair: '}',
		open: true,
	},
	'[': {
		pair: ']',
		open: true,
	},
	')': {
		pair: '(',
		open: false,
	},
	'}': {
		pair: '{',
		open: false,
	},
	']': {
		pair: '[',
		open: false,
	},
}

type stack []rune

func (s stack) push(r rune) stack {
	return append(s, r)
}

func (s stack) pop() (stack, *rune) {
	if len(s) == 0 {
		return s, nil
	}

	r := s[len(s)-1]
	s = s[:len(s)-1]
	return s, &r
}

func isValid(s string) bool {
	chars := stack{}
	for _, r := range s {
		bracket, ok := brackets[r]
		if !ok {
			return false
		}

		if bracket.open {
			chars = chars.push(r)
		} else {
			var prev *rune
			chars, prev = chars.pop()
			if prev == nil || *prev != bracket.pair {
				return false
			}
		}
	}

	return len(chars) == 0
}
