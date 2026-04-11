package insersection_of_two_arrays

import "testing"

func BenchmarkInsersectionOfTwoArrays(b *testing.B) {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InsersectionOfTwoArrays(nums1, nums2)
	}
}

func BenchmarkInsersectionOfTwoArraysLarge(b *testing.B) {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	nums2 := []int{8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InsersectionOfTwoArrays(nums1, nums2)
	}
}
