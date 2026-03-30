package palindrome_number

import "testing"

func TestPalindromeNumber(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{121, true},
		{-121, false},
		{1001, true},
		{10, false},
	}

	for _, tt := range tests {
		result := PalindromeNumber(tt.input)
		if result != tt.expected {
			t.Errorf("got %t, want %t", result, tt.expected)
		}
	}
}
