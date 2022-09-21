package main

import (
	"crypto/sha256"
	"testing"
)

func TestDiffBits(t *testing.T) {
	tests := []struct {
		x    [sha256.Size]byte
		y    [sha256.Size]byte
		want int
	}{
		{
			x:    [sha256.Size]byte{},
			y:    [sha256.Size]byte{},
			want: 0,
		},
		{
			x:    [sha256.Size]byte{1, 1, 1, 1, 1},
			y:    [sha256.Size]byte{},
			want: 5,
		},

		{
			x:    [sha256.Size]byte{1, 2, 3, 4, 5},
			y:    [sha256.Size]byte{1, 2, 4, 3, 5},
			want: 3,
		},
		{
			x:    sha256.Sum256([]byte("x")),
			y:    sha256.Sum256([]byte("X")),
			want: 44,
		},
	}
	for _, tt := range tests {
		if got := diffBits(tt.x, tt.y); got != tt.want {
			t.Errorf("diffBits(\n  %x\n  %x\n) = %d; want: %d\n",
				tt.x, tt.y, got, tt.want)
		}
	}
}
