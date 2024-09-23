package happy

import "strconv"

func isHappy(num int) bool {
	seen := map[int]struct{}{}
	for {
		if num == 1 {
			return true
		} else if _, ok := seen[num]; ok {
			return false
		}

		seen[num] = struct{}{}
		str := strconv.Itoa(num)
		cur := 0
		for _, s := range str {
			n, err := strconv.Atoi(string(s))
			if err != nil {
				return false
			}
			cur += n * n
		}
		num = cur
	}
}
