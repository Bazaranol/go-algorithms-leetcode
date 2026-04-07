package move_zeroes

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "basic case with zeros",
			nums: []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "no zeros",
			nums: []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "all zeros",
			nums: []int{0, 0, 0},
			want: []int{0, 0, 0},
		},
		{
			name: "single non-zero element",
			nums: []int{1},
			want: []int{1},
		},
		{
			name: "single zero element",
			nums: []int{0},
			want: []int{0},
		},
		{
			name: "zero at start",
			nums: []int{0, 1, 2},
			want: []int{1, 2, 0},
		},
		{
			name: "zero at end",
			nums: []int{1, 2, 0},
			want: []int{1, 2, 0},
		},
		{
			name: "multiple zeros scattered",
			nums: []int{0, 0, 1, 0, 2, 0, 3},
			want: []int{1, 2, 3, 0, 0, 0, 0},
		},
		{
			name: "zeros and non-zeros alternating",
			nums: []int{1, 0, 2, 0, 3, 0},
			want: []int{1, 2, 3, 0, 0, 0},
		},
		{
			name: "negative numbers",
			nums: []int{-1, 0, 2, 0, -3},
			want: []int{-1, 2, -3, 0, 0},
		},
		{
			name: "large numbers",
			nums: []int{1000, 0, 2000, 0, 3000},
			want: []int{1000, 2000, 3000, 0, 0},
		},
		{
			name: "two elements with zero",
			nums: []int{0, 1},
			want: []int{1, 0},
		},
		{
			name: "two elements without zero",
			nums: []int{1, 2},
			want: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy to avoid modifying original test data
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)

			got := MoveZeroes(numsCopy)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MoveZeroes(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
