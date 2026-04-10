package contains_duplicate

import "testing"

func BenchmarkContainsDuplicate(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsDuplicate(nums)
	}
}

func BenchmarkContainsDuplicateNoMatch(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsDuplicate(nums)
	}
}
