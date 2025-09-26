package captcha

import (
	"fmt"
	"testing"
)

func TestMin3_1(t *testing.T) {
	type test struct {
		name    string
		x, y, z uint8
		want    uint8
	}

	tests := []test{
		{"x is min", 1, 2, 3, 1},
		{"y is min", 3, 1, 2, 1},
		{"z is min", 2, 3, 1, 1},
		{"all equal", 1, 1, 1, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := min3(tt.x, tt.y, tt.z); got != tt.want {
				t.Errorf("min3() = %v, want %v", got, tt.want)
			}
		})
	}
}
