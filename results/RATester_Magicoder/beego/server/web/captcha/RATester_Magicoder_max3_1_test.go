package captcha

import (
	"fmt"
	"testing"
)

func Testmax3_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		x, y, z uint8
		want    uint8
	}

	tests := []test{
		{1, 2, 3, 3},
		{4, 5, 6, 6},
		{7, 8, 9, 9},
		{10, 11, 12, 12},
		{13, 14, 15, 15},
	}

	for _, tt := range tests {
		got := max3(tt.x, tt.y, tt.z)
		if got != tt.want {
			t.Errorf("max3(%d, %d, %d) = %d, want %d", tt.x, tt.y, tt.z, got, tt.want)
		}
	}
}
