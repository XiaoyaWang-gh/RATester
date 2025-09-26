package captcha

import (
	"fmt"
	"testing"
)

func TestMax3_1(t *testing.T) {
	type test struct {
		name    string
		x, y, z uint8
		want    uint8
	}

	tests := []test{
		{"x is max", 10, 5, 3, 10},
		{"y is max", 3, 10, 5, 10},
		{"z is max", 5, 3, 10, 10},
		{"all equal", 10, 10, 10, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := max3(tt.x, tt.y, tt.z); got != tt.want {
				t.Errorf("max3() = %v, want %v", got, tt.want)
			}
		})
	}
}
