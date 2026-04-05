package missing_number

import "testing"

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "missing 0 from [1]",
			nums: []int{1},
			want: 0,
		},
		{
			name: "missing 2 from [0, 1]",
			nums: []int{0, 1},
			want: 2,
		},
		{
			name: "missing 2 from [0, 1, 3]",
			nums: []int{0, 1, 3},
			want: 2,
		},
		{
			name: "missing 8 from [9, 6, 4, 2, 3, 5, 7, 0, 1]",
			nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			want: 8,
		},
		{
			name: "missing 0 from [1, 2, 3]",
			nums: []int{1, 2, 3},
			want: 0,
		},
		{
			name: "missing 1 from [0, 2, 3]",
			nums: []int{0, 2, 3},
			want: 1,
		},
		{
			name: "missing 3 from [0, 1, 2]",
			nums: []int{0, 1, 2},
			want: 3,
		},
		{
			name: "single element missing 0",
			nums: []int{1},
			want: 0,
		},
		{
			name: "single element missing 1",
			nums: []int{0},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MissingNumber(tt.nums)
			if got != tt.want {
				t.Errorf("MissingNumber(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
