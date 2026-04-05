package power_of_four

import "testing"

func BenchmarkPowerOfFour(b *testing.B) {
	n := 1024
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PowerOfFour(n)
	}
}

func BenchmarkPowerOfFourByBit(b *testing.B) {
	n := 1024
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PowerOfFourByBit(n)
	}
}
