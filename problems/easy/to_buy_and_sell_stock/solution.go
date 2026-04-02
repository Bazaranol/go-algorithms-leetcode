package to_buy_and_sell_stock

func ToBuyAndSellStock(prices []int) int {
	var min_price int = prices[0]
	var max_profit int = 0

	for _, value := range prices {

		if newProfit := value - min_price; newProfit > max_profit {
			max_profit = newProfit
		}

		if value < min_price {
			min_price = value
		}

	}

	return max_profit
}
