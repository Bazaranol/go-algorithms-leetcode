package missing_number

import "testing"

func BenchmarkMissingNumber(b *testing.B) {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MissingNumber(nums)
	}
}
