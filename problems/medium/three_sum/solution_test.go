package three_sum

import (
	"reflect"
	"slices"
	"testing"
)

func sortTriplets(triplets [][]int) {
	for _, t := range triplets {
		slices.Sort(t)
	}
	slices.SortFunc(triplets, func(a, b []int) int {
		for i := 0; i < 3; i++ {
			if a[i] < b[i] {
				return -1
			}
			if a[i] > b[i] {
				return 1
			}
		}
		return 0
	})
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "basic case with negative zero positive",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name: "single triplet",
			nums: []int{0, 0, 0},
			want: [][]int{
				{0, 0, 0},
			},
		},
		{
			name: "no triplets sum to zero",
			nums: []int{1, 2, 3},
			want: nil,
		},
		{
			name: "with duplicates",
			nums: []int{-2, 0, 1, 1, 2},
			want: [][]int{
				{-2, 0, 2},
				{-2, 1, 1},
			},
		},
		{
			name: "more complex duplicates",
			nums: []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6},
			want: [][]int{
				{-4, -2, 6},
				{-4, 0, 4},
				{-4, 1, 3},
				{-4, 2, 2},
				{-2, -2, 4},
				{-2, 0, 2},
			},
		},
		{
			name: "simple negative triplet",
			nums: []int{-1, -1, -1},
			want: nil,
		},
		{
			name: "with zeros",
			nums: []int{-4, 0, 0, 0, 1, 2, 3},
			want: [][]int{
				{-4, 1, 3},
				{0, 0, 0},
			},
		},
		{
			name: "ascending order range",
			nums: []int{-3, -2, -1, 0, 1, 2, 3},
			want: [][]int{
				{-3, 0, 3},
				{-3, 1, 2},
				{-2, -1, 3},
				{-2, 0, 2},
				{-1, 0, 1},
			},
		},
		{
			name: "mixed array with multiple solutions",
			nums: []int{-6, -4, -1, -1, 0, 1, 2, 3, 4, 5},
			want: [][]int{
				{-6, 1, 5},
				{-6, 2, 4},
				{-4, -1, 5},
				{-4, 0, 4},
				{-4, 1, 3},
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name: "no valid triplets with large numbers",
			nums: []int{1000, 2000, 3000, 4000},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ThreeSum(tt.nums)
			// Sort both for comparison since order doesn't matter
			sortTriplets(got)
			sortTriplets(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThreeSum(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
