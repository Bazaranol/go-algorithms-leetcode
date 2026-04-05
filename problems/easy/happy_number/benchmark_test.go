package happy_number

import "testing"

func BenchmarkHappyNumber(b *testing.B) {
	n := 1000
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HappyNumber(n)
	}
}
