package twosum

import "testing"

func BenchmarkTwoSum(b *testing.B) {
	nums := []int{2, 7, 11, 15}

	for i := 0; i < b.N; i++ {
		TwoSum(nums, 9)
	}
}
