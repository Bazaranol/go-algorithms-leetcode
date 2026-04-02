package to_buy_and_sell_stock

import "testing"

func TestToBuyAndSellStock(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "basic case",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			name:   "decreasing prices",
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		{
			name:   "single element",
			prices: []int{5},
			want:   0,
		},
		{
			name:   "two elements with profit",
			prices: []int{2, 4},
			want:   2,
		},
		{
			name:   "two elements no profit",
			prices: []int{4, 2},
			want:   0,
		},
		{
			name:   "increasing prices",
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
		{
			name:   "profit in middle",
			prices: []int{3, 2, 6, 5, 0, 3},
			want:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToBuyAndSellStock(tt.prices)
			if got != tt.want {
				t.Errorf("ToBuyAndSellStock(%v) = %d, want %d", tt.prices, got, tt.want)
			}
		})
	}
}
