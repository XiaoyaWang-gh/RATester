package template

import (
	"fmt"
	"testing"
)

func TestIsHexInt_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{"Test Case 1", "0x1234", true},
		{"Test Case 2", "0X1234", true},
		{"Test Case 3", "0x", false},
		{"Test Case 4", "1234", false},
		{"Test Case 5", "0p1234", false},
		{"Test Case 6", "0Xp1234", false},
		{"Test Case 7", "0x123g", false},
		{"Test Case 8", "0X123G", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isHexInt(tt.arg); got != tt.want {
				t.Errorf("isHexInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
