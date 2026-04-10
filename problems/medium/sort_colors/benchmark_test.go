package sort_colors

import "testing"

func BenchmarkSortColors(b *testing.B) {
	nums := []int{2, 0, 2, 1, 1, 0, 2, 1, 0, 2, 0, 1}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Make a copy for each iteration
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		SortColors(numsCopy)
	}
}
