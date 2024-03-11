package roman

var symbolTable map[rune]int = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

// fastest solution, combining a mapping of runes to numeric values and
// looking ahead to determine whether to subtract or add
func romanToInt(in string) int {
	num := 0
	for i, s := range in {
		add := symbolTable[s]
		if i < len(in)-1 && add < symbolTable[rune(in[i+1])] {
			num -= add
		} else {
			num += add
		}
	}
	return num
}

// runes -> rune that can be placed before them for subtraction
var subTable map[rune]rune = map[rune]rune{
	'V': 'I',
	'X': 'I',
	'L': 'X',
	'C': 'X',
	'D': 'C',
	'M': 'C',
}

func romanToIntSubTable(in string) int {
	num := 0
	var prev rune
	for _, s := range in {
		value := symbolTable[s]
		if prev != 0 {
			if sub, ok := subTable[s]; ok && sub == prev {
				subValue := symbolTable[sub]
				value = value - subValue - subValue
			}
		}
		prev = s
		num += value
	}
	return num
}
