package maximum_subarray

import "testing"

func TestMaximumSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "basic positive numbers",
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6, // [4, -1, 2, 1]
		},
		{
			name: "single element positive",
			nums: []int{5},
			want: 5,
		},
		{
			name: "single element negative",
			nums: []int{-3},
			want: -3,
		},
		{
			name: "all positive numbers",
			nums: []int{1, 2, 3, 4, 5},
			want: 15, // entire array
		},
		{
			name: "all negative numbers",
			nums: []int{-5, -2, -8, -1, -4},
			want: -1, // single largest element
		},
		{
			name: "mixed with zeros",
			nums: []int{-2, 0, -1, 2, 0, 3},
			want: 5, // [2, 0, 3]
		},
		{
			name: "alternating positive and negative",
			nums: []int{5, -3, 5},
			want: 7, // [5, -3, 5]
		},
		{
			name: "single large positive surrounded by negatives",
			nums: []int{-10, -5, 100, -5, -10},
			want: 100,
		},
		{
			name: "two positive separated by negative",
			nums: []int{3, -2, 5},
			want: 6, // [3, -2, 5]
		},
		{
			name: "duplicate numbers",
			nums: []int{1, 1, 1, 1},
			want: 4, // entire array
		},
		{
			name: "long array with peak in middle",
			nums: []int{-1, -2, -3, 10, 20, 30, -5, -10},
			want: 60, // [10, 20, 30]
		},
		{
			name: "two adjacent negatives",
			nums: []int{-1, -2},
			want: -1,
		},
		{
			name: "positive then large negative",
			nums: []int{5, 6, 7, -1000, 10},
			want: 18, // [5, 6, 7]
		},
		{
			name: "contains zero",
			nums: []int{-1, 0, -2},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaximumSubarray(tt.nums)
			if got != tt.want {
				t.Errorf("MaximumSubarray(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
