package template

import (
	"fmt"
	"testing"
)

func TestasciiAlpha_1(t *testing.T) {
	tests := []struct {
		name string
		c    byte
		want bool
	}{
		{"Test Case 1", 'A', true},
		{"Test Case 2", 'Z', true},
		{"Test Case 3", 'a', true},
		{"Test Case 4", 'z', true},
		{"Test Case 5", '0', false},
		{"Test Case 6", '9', false},
		{"Test Case 7", '@', false},
		{"Test Case 8", '[', false},
		{"Test Case 9", '`', false},
		{"Test Case 10", '{', false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := asciiAlpha(tt.c); got != tt.want {
				t.Errorf("asciiAlpha() = %v, want %v", got, tt.want)
			}
		})
	}
}
