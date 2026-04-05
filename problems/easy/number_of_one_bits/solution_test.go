package number_of_one_bits

import "testing"

func TestNumberOfOneBits(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "number 0",
			n:    0,
			want: 0,
		},
		{
			name: "number 1 (1 bit set)",
			n:    1,
			want: 1,
		},
		{
			name: "number 2 (10 binary)",
			n:    2,
			want: 1,
		},
		{
			name: "number 3 (11 binary)",
			n:    3,
			want: 2,
		},
		{
			name: "number 11 (1011 binary)",
			n:    11,
			want: 3,
		},
		{
			name: "number 15 (1111 binary)",
			n:    15,
			want: 4,
		},
		{
			name: "number 128 (10000000 binary)",
			n:    128,
			want: 1,
		},
		{
			name: "number 255 (11111111 binary)",
			n:    255,
			want: 8,
		},
		{
			name: "number 1024 (10000000000 binary)",
			n:    1024,
			want: 1,
		},
		{
			name: "number 2147483647 (all 1s in 31 bits)",
			n:    2147483647,
			want: 31,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NumberOfOneBits(tt.n)
			if got != tt.want {
				t.Errorf("NumberOfOneBits(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
