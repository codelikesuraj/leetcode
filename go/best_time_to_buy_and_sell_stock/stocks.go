package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	maxProfit := 0

	if len(prices) < 1 {
		return maxProfit
	}

	buyPrice := prices[0]

	for _, price := range prices {
		if price < buyPrice {
			buyPrice = price
		} else {
			profit := price - buyPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}
