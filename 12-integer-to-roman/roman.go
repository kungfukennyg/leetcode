package roman

import (
	"strings"
)

var values = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var romans = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

func intToRoman(num int) string {
	var sb strings.Builder
	for i, v := range values {
		if num == 0 {
			break
		}
		for num >= v {
			sb.WriteString(romans[i])
			num -= v
			if num == 0 {
				break
			}
		}
	}

	return sb.String()
}
