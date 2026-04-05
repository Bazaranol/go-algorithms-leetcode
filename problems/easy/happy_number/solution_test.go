package happy_number

import "testing"

func TestHappyNumber(t *testing.T) {
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
			name: "number 7",
			n:    7,
			want: true,
		},
		{
			name: "number 10",
			n:    10,
			want: true,
		},
		{
			name: "number 13",
			n:    13,
			want: true,
		},
		{
			name: "number 19",
			n:    19,
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
			name: "number 4",
			n:    4,
			want: false,
		},
		{
			name: "number 5",
			n:    5,
			want: false,
		},
		{
			name: "number 6",
			n:    6,
			want: false,
		},
		{
			name: "number 100",
			n:    100,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HappyNumber(tt.n)
			if got != tt.want {
				t.Errorf("HappyNumber(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
