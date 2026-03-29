package roman_to_integer

import "testing"

func BenchmarkRomanToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RomanToInt("MCMXCIV")
	}
}
