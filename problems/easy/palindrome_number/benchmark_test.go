package palindrome_number

import "testing"

func BenchmarkPalindromeNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PalindromeNumber(123321)
	}
}
