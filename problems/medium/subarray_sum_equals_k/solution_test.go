package subarray_sum_equals_k

import "testing"

func TestSubarraySumEqualsK(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "basic case",
			nums: []int{1, 1, 1},
			k:    2,
			want: 2, // [1,1] at (0,1) and [1,1] at (1,2)
		},
		{
			name: "single element matches",
			nums: []int{1, 2, 3},
			k:    3,
			want: 2, // [3] and [1,2]
		},
		{
			name: "no matches",
			nums: []int{1, 2, 3},
			k:    7,
			want: 0,
		},
		{
			name: "single element array match",
			nums: []int{5},
			k:    5,
			want: 1,
		},
		{
			name: "single element array no match",
			nums: []int{5},
			k:    3,
			want: 0,
		},
		{
			name: "entire array matches",
			nums: []int{1, 2, 3},
			k:    6,
			want: 1, // [1,2,3]
		},
		{
			name: "with negative numbers",
			nums: []int{1, 2, -2, 3},
			k:    3,
			want: 3, // [1,2], [2,-2,3], [3]
		},
		{
			name: "all zeros",
			nums: []int{0, 0, 0, 0},
			k:    0,
			want: 10, // 1 + 2 + 3 + 4 = 10 subarrays (all contiguous subarrays sum to 0)
		},
		{
			name: "multiple occurrences",
			nums: []int{1, 1, 1, 1, 1},
			k:    2,
			want: 4, // (0,1), (1,2), (2,3), (3,4)
		},
		{
			name: "k equals zero",
			nums: []int{-1, -1, 1, 1},
			k:    0,
			want: 2, // [-1,1] at (0,2) and [-1,1] at (1,3)
		},
		{
			name: "long array",
			nums: []int{1, 2, 1, 2, 1},
			k:    3,
			want: 4, // [1,2] at (0,1), [2,1] at (1,2), [1,2] at (2,3), [2,1] at (3,4)
		},
		{
			name: "duplicates same subarray",
			nums: []int{2, 1, 2, 1, 2},
			k:    3,
			want: 4, // [2,1] at (0,1), [1,2] at (1,2), [2,1] at (2,3), [1,2] at (3,4)
		},
		{
			name: "prefix sum technique",
			nums: []int{10, 5, 2, 7, 1, 9},
			k:    15,
			want: 2, // [10,5] and [5,2,7,1]
		},
		{
			name: "all negative",
			nums: []int{-2, -1, -3, -1, -2},
			k:    -3,
			want: 3, // [-2,-1], [-3], [-1,-2]
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SubarraySumEqualsK(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("SubarraySumEqualsK(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
