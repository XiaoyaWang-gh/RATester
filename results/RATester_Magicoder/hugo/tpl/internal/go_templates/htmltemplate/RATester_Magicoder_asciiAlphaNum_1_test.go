package template

import (
	"fmt"
	"testing"
)

func TestasciiAlphaNum_1(t *testing.T) {
	tests := []struct {
		name string
		c    byte
		want bool
	}{
		{"Test Case 1", 'a', true},
		{"Test Case 2", 'A', true},
		{"Test Case 3", '1', true},
		{"Test Case 4", ' ', false},
		{"Test Case 5", '@', false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := asciiAlphaNum(tt.c); got != tt.want {
				t.Errorf("asciiAlphaNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
