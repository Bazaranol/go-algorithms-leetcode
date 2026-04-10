package sort_colors

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "basic case",
			nums: []int{2, 0, 2, 1, 1, 0},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "already sorted",
			nums: []int{0, 1, 2},
			want: []int{0, 1, 2},
		},
		{
			name: "reverse sorted",
			nums: []int{2, 1, 0},
			want: []int{0, 1, 2},
		},
		{
			name: "all zeros",
			nums: []int{0, 0, 0},
			want: []int{0, 0, 0},
		},
		{
			name: "all ones",
			nums: []int{1, 1, 1},
			want: []int{1, 1, 1},
		},
		{
			name: "all twos",
			nums: []int{2, 2, 2},
			want: []int{2, 2, 2},
		},
		{
			name: "single element zero",
			nums: []int{0},
			want: []int{0},
		},
		{
			name: "single element one",
			nums: []int{1},
			want: []int{1},
		},
		{
			name: "single element two",
			nums: []int{2},
			want: []int{2},
		},
		{
			name: "two elements",
			nums: []int{1, 0},
			want: []int{0, 1},
		},
		{
			name: "twos at start",
			nums: []int{2, 2, 0, 1, 2},
			want: []int{0, 1, 2, 2, 2},
		},
		{
			name: "zeros at end",
			nums: []int{1, 2, 0, 0, 1},
			want: []int{0, 0, 1, 1, 2},
		},
		{
			name: "mixed with many duplicates",
			nums: []int{2, 0, 1, 2, 0, 1, 0, 2, 1, 0},
			want: []int{0, 0, 0, 0, 1, 1, 1, 2, 2, 2},
		},
		{
			name: "mostly zeros",
			nums: []int{0, 0, 0, 2, 1},
			want: []int{0, 0, 0, 1, 2},
		},
		{
			name: "mostly twos",
			nums: []int{2, 1, 0, 2, 2},
			want: []int{0, 1, 2, 2, 2},
		},
		{
			name: "alternating",
			nums: []int{0, 2, 1, 2, 0, 1},
			want: []int{0, 0, 1, 1, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy to avoid modifying original test data
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)

			got := SortColors(numsCopy)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortColors(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
