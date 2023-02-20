package main

import "testing"

func TestCanClimb2(t *testing.T) {
	tests := []struct {
		name string
		from uint8
		to   uint8
		want bool
	}{
		{
			name: "can go one up",
			from: 'a',
			to:   'b',
			want: true,
		},

		{
			name: "can not go two up",
			from: 'a',
			to:   'c',
			want: false,
		},
		{
			name: "can go down as much as i want",
			from: 'z',
			to:   'a',
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canClimb(tt.from, tt.to); got != tt.want {
				t.Errorf("CanClimb2() = %v, want %v", got, tt.want)
			}
		})
	}
}
