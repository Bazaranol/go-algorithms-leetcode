package container_with_most_water

import "testing"

func TestContainerWithMostWater(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "basic case",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49, // min(8, 7) * (8 - 1) = 7 * 7 = 49
		},
		{
			name:   "two elements",
			height: []int{1, 1},
			want:   1, // min(1, 1) * (1 - 0) = 1 * 1 = 1
		},
		{
			name:   "increasing heights",
			height: []int{1, 2, 3, 4, 5},
			want:   6, // min(2, 5) * (4 - 1) = 2 * 3 = 6
		},
		{
			name:   "decreasing heights",
			height: []int{5, 4, 3, 2, 1},
			want:   6, // min(5, 2) * (3 - 0) = 2 * 3 = 6
		},
		{
			name:   "all same height",
			height: []int{3, 3, 3, 3},
			want:   9, // min(3, 3) * (3 - 0) = 3 * 3 = 9
		},
		{
			name:   "one tall bar at start",
			height: []int{10, 1, 1, 1},
			want:   3, // min(10, 1) * (3 - 0) = 1 * 3 = 3
		},
		{
			name:   "one tall bar at end",
			height: []int{1, 1, 1, 10},
			want:   3, // min(1, 10) * (3 - 0) = 1 * 3 = 3
		},
		{
			name:   "tall bars at edges",
			height: []int{10, 1, 2, 3, 10},
			want:   40, // min(10, 10) * (4 - 0) = 10 * 4 = 40
		},
		{
			name:   "large heights with small width",
			height: []int{100, 50},
			want:   50, // min(100, 50) * (1 - 0) = 50 * 1 = 50
		},
		{
			name:   "mixed heights",
			height: []int{2, 3, 4, 5, 18, 17, 6},
			want:   17, // min(17, 6) * (5 - 1) = 6 * 4 = 24? or min(18, 17) * (4 - 0) = 17 * 4 = 68? Need to check
		},
		{
			name:   "zero heights",
			height: []int{0, 1, 0},
			want:   0, // min(0, 0) * (2 - 0) = 0 * 2 = 0
		},
		{
			name:   "three elements valley",
			height: []int{2, 1, 2},
			want:   4, // min(2, 2) * (2 - 0) = 2 * 2 = 4
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainerWithMostWater(tt.height)
			if got != tt.want {
				t.Errorf("ContainerWithMostWater(%v) = %d, want %d", tt.height, got, tt.want)
			}
		})
	}
}
