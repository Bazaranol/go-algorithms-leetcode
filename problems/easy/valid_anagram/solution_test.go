package valid_anagram

import "testing"

func TestValidAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "basic anagram",
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			name: "not anagram",
			s:    "rat",
			t:    "car",
			want: false,
		},
		{
			name: "single character",
			s:    "a",
			t:    "a",
			want: true,
		},
		{
			name: "single character different",
			s:    "a",
			t:    "b",
			want: false,
		},
		{
			name: "empty strings",
			s:    "",
			t:    "",
			want: true,
		},
		{
			name: "one empty string",
			s:    "a",
			t:    "",
			want: false,
		},
		{
			name: "different lengths",
			s:    "ab",
			t:    "abc",
			want: false,
		},
		{
			name: "same string",
			s:    "hello",
			t:    "hello",
			want: true,
		},
		{
			name: "reversed string",
			s:    "abc",
			t:    "cba",
			want: true,
		},
		{
			name: "with spaces",
			s:    "a b c",
			t:    "c a b",
			want: true,
		},
		{
			name: "with numbers",
			s:    "12321",
			t:    "31221",
			want: true,
		},
		{
			name: "repeated characters",
			s:    "aabbcc",
			t:    "abcabc",
			want: true,
		},
		{
			name: "repeated characters not anagram",
			s:    "aabbcc",
			t:    "abcabd",
			want: false,
		},
		{
			name: "unicode characters",
			s:    "café",
			t:    "féca",
			want: true,
		},
		{
			name: "long string",
			s:    "theeyes",
			t:    "theseye",
			want: true, // Same character counts
		},
		{
			name: "long anagram",
			s:    "listen",
			t:    "silent",
			want: true,
		},
		{
			name: "case sensitive",
			s:    "Hello",
			t:    "hello",
			want: false,
		},
		{
			name: "repeated in different order",
			s:    "aaa",
			t:    "aaa",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidAnagram(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("ValidAnagram(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
