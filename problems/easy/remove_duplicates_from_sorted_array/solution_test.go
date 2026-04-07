package remove_duplicates_from_sorted_array

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		want     int
		wantNums []int
	}{
		{
			name:     "array with duplicates",
			nums:     []int{1, 1, 2},
			want:     2,
			wantNums: []int{1, 2},
		},
		{
			name:     "array with multiple duplicates",
			nums:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want:     5,
			wantNums: []int{0, 1, 2, 3, 4},
		},
		{
			name:     "array without duplicates",
			nums:     []int{1, 2, 3, 4, 5},
			want:     5,
			wantNums: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "single element",
			nums:     []int{1},
			want:     1,
			wantNums: []int{1},
		},
		{
			name:     "two identical elements",
			nums:     []int{5, 5},
			want:     1,
			wantNums: []int{5},
		},
		{
			name:     "all elements same",
			nums:     []int{7, 7, 7, 7, 7},
			want:     1,
			wantNums: []int{7},
		},
		{
			name:     "duplicates at start",
			nums:     []int{1, 1, 1, 2, 3},
			want:     3,
			wantNums: []int{1, 2, 3},
		},
		{
			name:     "duplicates at end",
			nums:     []int{1, 2, 3, 3, 3},
			want:     3,
			wantNums: []int{1, 2, 3},
		},
		{
			name:     "duplicates in middle",
			nums:     []int{1, 2, 2, 2, 3},
			want:     3,
			wantNums: []int{1, 2, 3},
		},
		{
			name:     "negative numbers",
			nums:     []int{-5, -5, -2, 0, 0, 2, 2, 3},
			want:     5,
			wantNums: []int{-5, -2, 0, 2, 3},
		},
		{
			name:     "large array with minimal duplicates",
			nums:     []int{1, 2, 3, 4, 4, 5, 6, 7, 8, 9},
			want:     9,
			wantNums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy to avoid modifying original test data
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)

			got := RemoveDuplicatesFromSortedArray(numsCopy)

			if got != tt.want {
				t.Errorf("RemoveDuplicatesFromSortedArray() returned %d, want %d", got, tt.want)
			}

			// Check first 'got' elements match expected
			if !reflect.DeepEqual(numsCopy[:got], tt.wantNums) {
				t.Errorf("First %d elements are %v, want %v", got, numsCopy[:got], tt.wantNums)
			}
		})
	}
}
