package container_with_most_water

import "testing"

func BenchmarkContainerWithMostWater(b *testing.B) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7, 2, 5, 5, 8, 4, 7}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainerWithMostWater(height)
	}
}
