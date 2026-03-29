package roman_to_integer

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for _, tt := range tests {
		result := RomanToInt(tt.input)
		if result != tt.expected {
			t.Errorf("got %d, want %d", result, tt.expected)
		}
	}
}
