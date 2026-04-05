package power_of_four

import "testing"

func TestPowerOfFour(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{
			name: "number 1",
			n:    1,
			want: true,
		},
		{
			name: "number 4",
			n:    4,
			want: true,
		},
		{
			name: "number 16",
			n:    16,
			want: true,
		},
		{
			name: "number 64",
			n:    64,
			want: true,
		},
		{
			name: "number 256",
			n:    256,
			want: true,
		},
		{
			name: "number 1024",
			n:    1024,
			want: true,
		},
		{
			name: "number 2",
			n:    2,
			want: false,
		},
		{
			name: "number 3",
			n:    3,
			want: false,
		},
		{
			name: "number 5",
			n:    5,
			want: false,
		},
		{
			name: "number 8 (power of 2 but not 4)",
			n:    8,
			want: false,
		},
		{
			name: "number 32 (power of 2 but not 4)",
			n:    32,
			want: false,
		},
		{
			name: "number 0",
			n:    0,
			want: false,
		},
		{
			name: "negative number",
			n:    -1,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PowerOfFour(tt.n)
			if got != tt.want {
				t.Errorf("PowerOfFour(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func TestPowerOfFourByBit(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{
			name: "number 1",
			n:    1,
			want: true,
		},
		{
			name: "number 4",
			n:    4,
			want: true,
		},
		{
			name: "number 16",
			n:    16,
			want: true,
		},
		{
			name: "number 64",
			n:    64,
			want: true,
		},
		{
			name: "number 256",
			n:    256,
			want: true,
		},
		{
			name: "number 1024",
			n:    1024,
			want: true,
		},
		{
			name: "number 2",
			n:    2,
			want: false,
		},
		{
			name: "number 3",
			n:    3,
			want: false,
		},
		{
			name: "number 5",
			n:    5,
			want: false,
		},
		{
			name: "number 8 (power of 2 but not 4)",
			n:    8,
			want: false,
		},
		{
			name: "number 32 (power of 2 but not 4)",
			n:    32,
			want: false,
		},
		{
			name: "number 0",
			n:    0,
			want: false,
		},
		{
			name: "negative number",
			n:    -1,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PowerOfFourByBit(tt.n)
			if got != tt.want {
				t.Errorf("PowerOfFourByBit(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
