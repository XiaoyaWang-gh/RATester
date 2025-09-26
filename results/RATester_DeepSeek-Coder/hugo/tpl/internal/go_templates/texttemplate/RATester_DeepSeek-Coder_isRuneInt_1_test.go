package template

import (
	"fmt"
	"testing"
)

func TestIsRuneInt_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{"Test Case 1", "'123", true},
		{"Test Case 2", "123", false},
		{"Test Case 3", "'", false},
		{"Test Case 4", "''", false},
		{"Test Case 5", "''123", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isRuneInt(tt.arg); got != tt.want {
				t.Errorf("isRuneInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
