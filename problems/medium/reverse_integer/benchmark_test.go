package reverse_integer

import "testing"

func BenchmarkReverseInteger(b *testing.B) {
	x := 1534236468
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReverseInteger(x)
	}
}
