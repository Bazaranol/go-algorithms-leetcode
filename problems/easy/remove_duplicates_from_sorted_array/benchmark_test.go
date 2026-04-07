package remove_duplicates_from_sorted_array

import "testing"

func BenchmarkRemoveDuplicatesFromSortedArray(b *testing.B) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Make a copy for each iteration
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		RemoveDuplicatesFromSortedArray(numsCopy)
	}
}
