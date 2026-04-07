package three_sum

import "testing"

func BenchmarkThreeSum(b *testing.B) {
	nums := []int{-1, 0, 1, 2, -1, -4, 3, -2, 4, 5, -3, 6, -5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Make a copy for each iteration
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		ThreeSum(numsCopy)
	}
}
