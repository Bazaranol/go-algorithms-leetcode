package contains_duplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "basic duplicate",
			nums: []int{1, 2, 3, 1},
			want: true,
		},
		{
			name: "no duplicates",
			nums: []int{1, 2, 3, 4},
			want: false,
		},
		{
			name: "duplicate at end",
			nums: []int{1, 2, 3, 4, 3},
			want: true,
		},
		{
			name: "duplicate at start",
			nums: []int{1, 1, 2, 3},
			want: true,
		},
		{
			name: "consecutive duplicates",
			nums: []int{1, 2, 2, 3},
			want: true,
		},
		{
			name: "single element",
			nums: []int{1},
			want: false,
		},
		{
			name: "two same elements",
			nums: []int{1, 1},
			want: true,
		},
		{
			name: "two different elements",
			nums: []int{1, 2},
			want: false,
		},
		{
			name: "all duplicates",
			nums: []int{5, 5, 5, 5, 5},
			want: true,
		},
		{
			name: "negative numbers with duplicate",
			nums: []int{-1, -2, -3, -1},
			want: true,
		},
		{
			name: "negative numbers no duplicate",
			nums: []int{-1, -2, -3, -4},
			want: false,
		},
		{
			name: "mixed positive and negative with duplicate",
			nums: []int{1, -1, 0, -1},
			want: true,
		},
		{
			name: "mixed positive and negative no duplicate",
			nums: []int{1, -1, 0, 2},
			want: false,
		},
		{
			name: "zero repeated",
			nums: []int{0, 0},
			want: true,
		},
		{
			name: "zero and other elements",
			nums: []int{1, 2, 3, 0, 4, 0},
			want: true,
		},
		{
			name: "large numbers with duplicate",
			nums: []int{1000000, 999999, 1000000},
			want: true,
		},
		{
			name: "large numbers no duplicate",
			nums: []int{1000000, 999999, 999998},
			want: false,
		},
		{
			name: "long array with duplicate at end",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1},
			want: true,
		},
		{
			name: "long array no duplicate",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsDuplicate(tt.nums)
			if got != tt.want {
				t.Errorf("ContainsDuplicate(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
