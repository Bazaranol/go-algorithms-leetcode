package count_primes

import "testing"

func BenchmarkCountPrimes(b *testing.B) {
	n := 10000
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountPrimes(n)
	}
}
