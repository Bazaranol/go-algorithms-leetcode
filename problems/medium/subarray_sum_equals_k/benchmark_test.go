package subarray_sum_equals_k

import "testing"

func BenchmarkSubarraySumEqualsK(b *testing.B) {
	nums := []int{1, 1, 1, 2, 1, 1, 1, 3, 2, 1, 2, 1}
	k := 2
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SubarraySumEqualsK(nums, k)
	}
}
