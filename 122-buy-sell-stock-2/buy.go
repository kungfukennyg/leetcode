package buysell2

import "math"

func maxProfit(prices []int) int {
	prevDay, profit := math.MinInt, 0
	for _, price := range prices {
		if prevDay == math.MinInt {
			prevDay = price
			continue
		}

		if price > prevDay {
			profit += price - prevDay
		}
		prevDay = price
	}
	return profit
}
