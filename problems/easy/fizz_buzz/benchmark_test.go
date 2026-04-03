package fizz_buzz

import "testing"

func BenchmarkFizzBuzz(b *testing.B) {
	n := 100
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FizzBuzz(n)
	}
}
