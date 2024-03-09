package lexico

func lexicalOrder(n int) []int {
	seq := make([]int, n)
	cur := 1
	for i := 1; i <= n; i++ {
		seq[i-1] = cur
		if cur*10 <= n {
			cur *= 10
		} else {
			for cur%10 == 9 || cur+1 > n {
				cur /= 10
			}
			cur += 1
		}
	}

	return seq
}
