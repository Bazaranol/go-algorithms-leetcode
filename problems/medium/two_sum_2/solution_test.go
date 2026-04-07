package two_sum_2

import (
	"reflect"
	"testing"
)

func TestTwoSum2(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{
			name:    "basic case",
			numbers: []int{2, 7, 11, 15},
			target:  9,
			want:    []int{1, 2}, // 2 + 7 = 9
		},
		{
			name:    "two elements",
			numbers: []int{1, 2},
			target:  3,
			want:    []int{1, 2},
		},
		{
			name:    "numbers at different positions",
			numbers: []int{1, 3, 4, 5, 7, 11},
			target:  9,
			want:    []int{3, 4}, // 4 + 5 = 9
		},
		{
			name:    "sum of first and last",
			numbers: []int{2, 3, 4},
			target:  6,
			want:    []int{1, 3}, // 2 + 4 = 6
		},
		{
			name:    "sorted array middle elements",
			numbers: []int{1, 2, 3, 4, 5, 6, 7},
			target:  11,
			want:    []int{4, 7}, // 4 + 7 = 11
		},
		{
			name:    "larger sorted array",
			numbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target:  17,
			want:    []int{7, 10}, // 8 + 9 = 17
		},
		{
			name:    "consecutive numbers",
			numbers: []int{1, 2, 3, 4, 5},
			target:  5,
			want:    []int{1, 4}, // 1 + 4 = 5
		},
		{
			name:    "duplicates in sorted array",
			numbers: []int{1, 1, 1, 2, 2, 3},
			target:  4,
			want:    []int{1, 6}, // 1 + 3 = 4
		},
		{
			name:    "all small positives",
			numbers: []int{5, 6, 7, 8, 9},
			target:  14,
			want:    []int{1, 5}, // 5 + 9 = 14
		},
		{
			name:    "negative and positive numbers",
			numbers: []int{-10, -5, 0, 5, 10},
			target:  0,
			want:    []int{1, 5}, // -10 + 10 = 0
		},
		{
			name:    "negative numbers only",
			numbers: []int{-15, -10, -5, -2},
			target:  -17,
			want:    []int{1, 4}, // -15 + (-2) = -17
		},
		{
			name:    "large sorted array with target in middle",
			numbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			target:  24,
			want:    []int{9, 15}, // 9 + 15 = 24
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum2(tt.numbers, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum2(%v, %d) = %v, want %v", tt.numbers, tt.target, got, tt.want)
			}
		})
	}
}
