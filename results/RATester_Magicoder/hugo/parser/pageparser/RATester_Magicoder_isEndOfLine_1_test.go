package pageparser

import (
	"fmt"
	"testing"
)

func TestisEndOfLine_1(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{"Test case 1", '\r', true},
		{"Test case 2", '\n', true},
		{"Test case 3", 'a', false},
		{"Test case 4", ' ', false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isEndOfLine(tt.r); got != tt.want {
				t.Errorf("isEndOfLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
