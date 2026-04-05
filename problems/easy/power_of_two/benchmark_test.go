package power_of_two

import "testing"

func BenchmarkPowerOfTwo(b *testing.B) {
	n := 1024
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PowerOfTwo(n)
	}
}

func BenchmarkPowerOfTwoByBit(b *testing.B) {
	n := 1024
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PowerOfTwoByBit(n)
	}
}
