package related

import (
	"fmt"
	"testing"
)

func Testnorm_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		num  int
		min  int
		max  int
		want int
	}

	tests := []test{
		{num: 10, min: 0, max: 100, want: 10},
		{num: 50, min: 0, max: 100, want: 50},
		{num: 100, min: 0, max: 100, want: 100},
		{num: 0, min: 0, max: 100, want: 0},
		{num: -10, min: 0, max: 100, want: 0},
		{num: 110, min: 0, max: 100, want: 100},
		{num: 50, min: 50, max: 50, want: 100},
		{num: 50, min: 0, max: 50, want: 100},
		{num: 0, min: 0, max: 0, want: 0},
	}

	for _, tt := range tests {
		got := norm(tt.num, tt.min, tt.max)
		if got != tt.want {
			t.Errorf("norm(%d, %d, %d) = %d, want %d", tt.num, tt.min, tt.max, got, tt.want)
		}
	}
}
