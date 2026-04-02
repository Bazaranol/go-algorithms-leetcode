package reverse_integer

import "testing"

func TestReverseInteger(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{
			name: "positive number",
			x:    123,
			want: 321,
		},
		{
			name: "negative number",
			x:    -123,
			want: -321,
		},
		{
			name: "number with trailing zeros",
			x:    120,
			want: 21,
		},
		{
			name: "single digit",
			x:    5,
			want: 5,
		},
		{
			name: "zero",
			x:    0,
			want: 0,
		},
		{
			name: "overflow positive",
			x:    1534236469,
			want: 0,
		},
		{
			name: "large positive number",
			x:    1534236468,
			want: 0,
		},
		{
			name: "large negative number",
			x:    -2147483648,
			want: 0,
		},
		{
			name: "simple two digit",
			x:    10,
			want: 1,
		},
		{
			name: "palindrome number",
			x:    121,
			want: 121,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseInteger(tt.x)
			if got != tt.want {
				t.Errorf("ReverseInteger(%d) = %d, want %d", tt.x, got, tt.want)
			}
		})
	}
}
