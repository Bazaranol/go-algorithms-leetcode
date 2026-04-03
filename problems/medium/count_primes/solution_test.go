package count_primes

import "testing"

func TestCountPrimes(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "n less than 2",
			n:    0,
			want: 0,
		},
		{
			name: "n equals 1",
			n:    1,
			want: 0,
		},
		{
			name: "n equals 2",
			n:    2,
			want: 0,
		},
		{
			name: "n equals 3",
			n:    3,
			want: 1,
		},
		{
			name: "n equals 10",
			n:    10,
			want: 4,
		},
		{
			name: "n equals 15",
			n:    15,
			want: 6,
		},
		{
			name: "n equals 20",
			n:    20,
			want: 8,
		},
		{
			name: "n equals 100",
			n:    100,
			want: 25,
		},
		{
			name: "n equals 1000",
			n:    1000,
			want: 168,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountPrimes(tt.n)
			if got != tt.want {
				t.Errorf("CountPrimes(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
