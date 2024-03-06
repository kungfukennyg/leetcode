package buysell

import "math"

func maxProfit(prices []int) int {
	cheapest := math.MinInt
	profit := 0
	for _, price := range prices {
		if cheapest == math.MinInt {
			cheapest = price
			continue
		}
		if price < cheapest {
			cheapest = price
			continue
		}
		if price-cheapest > profit {
			profit = price - cheapest
		}
	}
	if profit < 0 {
		return 0
	}

	return profit
}
