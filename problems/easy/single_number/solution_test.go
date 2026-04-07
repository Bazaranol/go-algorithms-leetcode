package single_number

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "single element",
			nums: []int{1},
			want: 1,
		},
		{
			name: "1 appears once, 2 appears twice",
			nums: []int{2, 2, 1},
			want: 1,
		},
		{
			name: "4 appears once, others appear twice",
			nums: []int{4, 1, 2, 1, 2},
			want: 4,
		},
		{
			name: "0 is single number",
			nums: []int{0, 1, 0, 1, 2},
			want: 2,
		},
		{
			name: "negative single number",
			nums: []int{-1, -1, 3, 3, 5},
			want: 5,
		},
		{
			name: "single from start",
			nums: []int{7, 1, 1, 2, 2, 3, 3},
			want: 7,
		},
		{
			name: "single from middle",
			nums: []int{1, 1, 2, 3, 3, 4, 4},
			want: 2,
		},
		{
			name: "single from end",
			nums: []int{5, 5, 1, 1, 2, 2, 3},
			want: 3,
		},
		{
			name: "large numbers",
			nums: []int{1000000, 999999, 1000000, 999999, 42},
			want: 42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SingleNumber(tt.nums)
			if got != tt.want {
				t.Errorf("SingleNumber(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestSingleNumberByXOR(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "single element",
			nums: []int{1},
			want: 1,
		},
		{
			name: "1 appears once, 2 appears twice",
			nums: []int{2, 2, 1},
			want: 1,
		},
		{
			name: "4 appears once, others appear twice",
			nums: []int{4, 1, 2, 1, 2},
			want: 4,
		},
		{
			name: "0 is single number",
			nums: []int{0, 1, 0, 1, 2},
			want: 2,
		},
		{
			name: "negative single number",
			nums: []int{-1, -1, 3, 3, 5},
			want: 5,
		},
		{
			name: "single from start",
			nums: []int{7, 1, 1, 2, 2, 3, 3},
			want: 7,
		},
		{
			name: "single from middle",
			nums: []int{1, 1, 2, 3, 3, 4, 4},
			want: 2,
		},
		{
			name: "single from end",
			nums: []int{5, 5, 1, 1, 2, 2, 3},
			want: 3,
		},
		{
			name: "large numbers",
			nums: []int{1000000, 999999, 1000000, 999999, 42},
			want: 42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SingleNumberByXOR(tt.nums)
			if got != tt.want {
				t.Errorf("SingleNumberByXOR(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
