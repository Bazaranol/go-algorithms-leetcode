package to_buy_and_sell_stock

import "testing"

func BenchmarkToBuyAndSellStock(b *testing.B) {
	prices := []int{7, 1, 5, 3, 6, 4, 2, 8, 9, 1, 2, 3}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ToBuyAndSellStock(prices)
	}
}
