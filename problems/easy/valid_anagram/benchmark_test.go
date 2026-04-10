package valid_anagram

import "testing"

func BenchmarkValidAnagram(b *testing.B) {
	s := "listen"
	t := "silent"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ValidAnagram(s, t)
	}
}
