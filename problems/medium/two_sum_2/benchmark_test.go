package two_sum_2

import "testing"

func BenchmarkTwoSum2(b *testing.B) {
	numbers := []int{2, 7, 11, 15, 20, 25, 30, 35, 40, 45}
	target := 30
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum2(numbers, target)
	}
}
