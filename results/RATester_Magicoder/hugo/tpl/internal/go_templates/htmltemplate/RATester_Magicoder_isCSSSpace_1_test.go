package template

import (
	"fmt"
	"testing"
)

func TestisCSSSpace_1(t *testing.T) {
	tests := []struct {
		name string
		b    byte
		want bool
	}{
		{"Test case 1", '\t', true},
		{"Test case 2", '\n', true},
		{"Test case 3", '\f', true},
		{"Test case 4", '\r', true},
		{"Test case 5", ' ', true},
		{"Test case 6", 'a', false},
		{"Test case 7", '1', false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isCSSSpace(tt.b); got != tt.want {
				t.Errorf("isCSSSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
