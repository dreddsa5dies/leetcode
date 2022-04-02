/*
Runtime: 112 ms, faster than 89.64% of Go online submissions for Best Time to Buy and Sell Stock.
Memory Usage: 8 MB, less than 90.23% of Go online submissions for Best Time to Buy and Sell Stock.
*/
package main

import "math"

func main() {}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	sell := prices[0]
	profit := 0.0
	for i := range prices {
		if prices[i] < sell {
			sell = prices[i]
		}
		profit = math.Max(profit, float64(prices[i]-sell))
	}
	return int(profit)
}
