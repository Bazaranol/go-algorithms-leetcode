package first_unique_character_in_a_string

import "testing"

func BenchmarkFirstUniqueCharacterInAString(b *testing.B) {
	s := "loveleetcode"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FirstUniqueCharacterInAString(s)
	}
}

func BenchmarkFirstUniqueCharacterInAStringLong(b *testing.B) {
	s := "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzz"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FirstUniqueCharacterInAString(s)
	}
}
