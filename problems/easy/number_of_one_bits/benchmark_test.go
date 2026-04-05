package number_of_one_bits

import "testing"

func BenchmarkNumberOfOneBits(b *testing.B) {
	n := 2147483647
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NumberOfOneBits(n)
	}
}
