package power_of_three

import "testing"

func TestPowerOfThree(t *testing.T) {
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
			name: "number 3",
			n:    3,
			want: true,
		},
		{
			name: "number 9",
			n:    9,
			want: true,
		},
		{
			name: "number 27",
			n:    27,
			want: true,
		},
		{
			name: "number 81",
			n:    81,
			want: true,
		},
		{
			name: "number 243",
			n:    243,
			want: true,
		},
		{
			name: "number 2187",
			n:    2187,
			want: true,
		},
		{
			name: "number 2",
			n:    2,
			want: false,
		},
		{
			name: "number 4",
			n:    4,
			want: false,
		},
		{
			name: "number 10",
			n:    10,
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
			got := PowerOfThree(tt.n)
			if got != tt.want {
				t.Errorf("PowerOfThree(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func TestPowerOfThreeByBit(t *testing.T) {
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
			name: "number 3",
			n:    3,
			want: true,
		},
		{
			name: "number 9",
			n:    9,
			want: true,
		},
		{
			name: "number 27",
			n:    27,
			want: true,
		},
		{
			name: "number 81",
			n:    81,
			want: true,
		},
		{
			name: "number 243",
			n:    243,
			want: true,
		},
		{
			name: "number 2187",
			n:    2187,
			want: true,
		},
		{
			name: "number 2",
			n:    2,
			want: false,
		},
		{
			name: "number 4",
			n:    4,
			want: false,
		},
		{
			name: "number 10",
			n:    10,
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
			got := PowerOfThreeByBit(tt.n)
			if got != tt.want {
				t.Errorf("PowerOfThreeByBit(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
