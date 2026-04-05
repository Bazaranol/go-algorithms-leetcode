package power_of_three

import "testing"

func BenchmarkPowerOfThree(b *testing.B) {
	n := 2187
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PowerOfThree(n)
	}
}

func BenchmarkPowerOfThreeByBit(b *testing.B) {
	n := 2187
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PowerOfThreeByBit(n)
	}
}
