package template

import (
	"fmt"
	"testing"
)

func TestisHex_1(t *testing.T) {
	tests := []struct {
		name  string
		input byte
		want  bool
	}{
		{"Test Case 1", '0', true},
		{"Test Case 2", '9', true},
		{"Test Case 3", 'a', true},
		{"Test Case 4", 'f', true},
		{"Test Case 5", 'A', true},
		{"Test Case 6", 'F', true},
		{"Test Case 7", 'g', false},
		{"Test Case 8", 'z', false},
		{"Test Case 9", ' ', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isHex(tt.input); got != tt.want {
				t.Errorf("isHex(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
