package template

import (
	"fmt"
	"testing"
)

func TestisHexInt_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{"Test Case 1", "0x123", true},
		{"Test Case 2", "0X123", true},
		{"Test Case 3", "0x", false},
		{"Test Case 4", "123", false},
		{"Test Case 5", "0x123p4", false},
		{"Test Case 6", "0X123P4", false},
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
