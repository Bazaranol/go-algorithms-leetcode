package single_number

import "testing"

func BenchmarkSingleNumber(b *testing.B) {
	nums := []int{4, 1, 2, 1, 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SingleNumber(nums)
	}
}

func BenchmarkSingleNumberByXOR(b *testing.B) {
	nums := []int{4, 1, 2, 1, 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SingleNumberByXOR(nums)
	}
}
