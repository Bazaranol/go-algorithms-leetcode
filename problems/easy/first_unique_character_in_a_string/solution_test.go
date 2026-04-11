package first_unique_character_in_a_string

import "testing"

func TestFirstUniqueCharacterInAString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "basic unique at start",
			s:    "leetcode",
			want: 0,
		},
		{
			name: "basic unique in middle",
			s:    "loveleetcode",
			want: 2,
		},
		{
			name: "no unique character",
			s:    "aabb",
			want: -1,
		},
		{
			name: "single character",
			s:    "a",
			want: 0,
		},
		{
			name: "two same characters",
			s:    "aa",
			want: -1,
		},
		{
			name: "two different characters",
			s:    "ab",
			want: 0,
		},
		{
			name: "unique at end",
			s:    "aabbc",
			want: 4,
		},
		{
			name: "all same characters",
			s:    "aaaa",
			want: -1,
		},
		{
			name: "with numbers",
			s:    "112233",
			want: -1,
		},
		{
			name: "with spaces",
			s:    "a a b b c",
			want: 8,
		},
		{
			name: "case sensitive",
			s:    "Aa",
			want: 0,
		},
		{
			name: "long string unique at end",
			s:    "abcdefghijklmnopqrstuvwxyza",
			want: 1,
		},
		{
			name: "long string no unique",
			s:    "aabbccddeeffgghhiijjkkllmm",
			want: -1,
		},
		{
			name: "unicode characters",
			s:    "café",
			want: 0,
		},
		{
			name: "repeated unicode",
			s:    "caféfé",
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FirstUniqueCharacterInAString(tt.s)
			if got != tt.want {
				t.Errorf("FirstUniqueCharacterInAString(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
