package move_zeroes

import "testing"

func BenchmarkMoveZeroes(b *testing.B) {
	nums := []int{0, 1, 0, 3, 12, 0, 0, 2, 0, 4}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Make a copy for each iteration
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		MoveZeroes(numsCopy)
	}
}
