package insersection_of_two_arrays

import (
	"reflect"
	"testing"
)

func TestInsersectionOfTwoArrays(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{
			name:  "basic intersection",
			nums1: []int{1, 2, 2, 1},
			nums2: []int{2, 2},
			want:  []int{2},
		},
		{
			name:  "intersection with multiple elements",
			nums1: []int{4, 9, 5},
			nums2: []int{9, 4, 9, 8, 4},
			want:  []int{9, 4},
		},
		{
			name:  "no intersection",
			nums1: []int{1, 2, 3},
			nums2: []int{4, 5, 6},
			want:  nil,
		},
		{
			name:  "partial intersection",
			nums1: []int{1, 2, 3, 4},
			nums2: []int{2, 4, 6, 8},
			want:  []int{2, 4},
		},
		{
			name:  "single common element",
			nums1: []int{1, 2, 3},
			nums2: []int{3},
			want:  []int{3},
		},
		{
			name:  "single common element reverse",
			nums1: []int{1},
			nums2: []int{1, 2, 3},
			want:  []int{1},
		},
		{
			name:  "all same elements",
			nums1: []int{5, 5, 5},
			nums2: []int{5, 5, 5},
			want:  []int{5},
		},
		{
			name:  "duplicates in nums2",
			nums1: []int{1, 2, 3},
			nums2: []int{2, 2, 2},
			want:  []int{2},
		},
		{
			name:  "empty first array",
			nums1: []int{},
			nums2: []int{1, 2, 3},
			want:  nil,
		},
		{
			name:  "empty second array",
			nums1: []int{1, 2, 3},
			nums2: []int{},
			want:  nil,
		},
		{
			name:  "both empty arrays",
			nums1: []int{},
			nums2: []int{},
			want:  nil,
		},
		{
			name:  "single element in each",
			nums1: []int{1},
			nums2: []int{1},
			want:  []int{1},
		},
		{
			name:  "single different elements",
			nums1: []int{1},
			nums2: []int{2},
			want:  nil,
		},
		{
			name:  "negative numbers with intersection",
			nums1: []int{-1, -2, -3},
			nums2: []int{-2, -3, -4},
			want:  []int{-2, -3},
		},
		{
			name:  "negative and positive intersection",
			nums1: []int{-1, 0, 1, 2},
			nums2: []int{-2, -1, 0, 2},
			want:  []int{-1, 0, 2},
		},
		{
			name:  "large numbers intersection",
			nums1: []int{1000000, 999999, 999998},
			nums2: []int{999999, 1000000, 1000001},
			want:  []int{999999, 1000000},
		},
		{
			name:  "preserves order from nums2",
			nums1: []int{1, 2, 3, 4, 5},
			nums2: []int{5, 3, 1},
			want:  []int{5, 3, 1},
		},
		{
			name:  "eliminates duplicates",
			nums1: []int{1, 1, 2, 2, 3, 3},
			nums2: []int{2, 2, 2, 3, 3},
			want:  []int{2, 3},
		},
		{
			name:  "long arrays with few common elements",
			nums1: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			nums2: []int{5, 6, 7},
			want:  []int{5, 6, 7},
		},
		{
			name:  "repeated pattern",
			nums1: []int{1, 2, 1, 2, 1, 2},
			nums2: []int{2, 1, 2, 1},
			want:  []int{2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InsersectionOfTwoArrays(tt.nums1, tt.nums2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsersectionOfTwoArrays(%v, %v) = %v, want %v", tt.nums1, tt.nums2, got, tt.want)
			}
		})
	}
}
