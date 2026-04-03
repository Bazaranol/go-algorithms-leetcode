package fizz_buzz

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{
			name: "empty array",
			n:    0,
			want: []string{},
		},
		{
			name: "single element",
			n:    1,
			want: []string{"1"},
		},
		{
			name: "fizz and buzz",
			n:    5,
			want: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			name: "fizzbuzz",
			n:    15,
			want: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
		{
			name: "up to 10",
			n:    10,
			want: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz"},
		},
		{
			name: "only numbers",
			n:    2,
			want: []string{"1", "2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FizzBuzz(tt.n)
			if len(got) != len(tt.want) {
				t.Errorf("FizzBuzz(%d) returned array of length %d, want %d", tt.n, len(got), len(tt.want))
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("FizzBuzz(%d)[%d] = %s, want %s", tt.n, i, got[i], tt.want[i])
				}
			}
		})
	}
}
