package maximum_subarray

import "testing"

func BenchmarkMaximumSubarray(b *testing.B) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaximumSubarray(nums)
	}
}
