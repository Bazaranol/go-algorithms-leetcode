package power_of_two

import "testing"

func TestPowerOfTwo(t *testing.T) {
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
			name: "number 2",
			n:    2,
			want: true,
		},
		{
			name: "number 3",
			n:    3,
			want: false,
		},
		{
			name: "number 4",
			n:    4,
			want: true,
		},
		{
			name: "number 5",
			n:    5,
			want: false,
		},
		{
			name: "number 8",
			n:    8,
			want: true,
		},
		{
			name: "number 16",
			n:    16,
			want: true,
		},
		{
			name: "number 1024",
			n:    1024,
			want: true,
		},
		{
			name: "number 1023",
			n:    1023,
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
			got := PowerOfTwo(tt.n)
			if got != tt.want {
				t.Errorf("PowerOfTwo(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func TestPowerOfTwoByBit(t *testing.T) {
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
			name: "number 2",
			n:    2,
			want: true,
		},
		{
			name: "number 3",
			n:    3,
			want: false,
		},
		{
			name: "number 4",
			n:    4,
			want: true,
		},
		{
			name: "number 5",
			n:    5,
			want: false,
		},
		{
			name: "number 8",
			n:    8,
			want: true,
		},
		{
			name: "number 16",
			n:    16,
			want: true,
		},
		{
			name: "number 1024",
			n:    1024,
			want: true,
		},
		{
			name: "number 1023",
			n:    1023,
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
			got := PowerOfTwoByBit(tt.n)
			if got != tt.want {
				t.Errorf("PowerOfTwoByBit(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
