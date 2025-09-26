package helpers

import (
	"fmt"
	"testing"
)

func TestFirstUpper_1(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"empty string", "", ""},
		{"single character", "a", "A"},
		{"multiple characters", "hello", "H"},
		{"multiple characters with spaces", "hello world", "H"},
		{"multiple characters with special characters", "hello world!", "H"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := FirstUpper(tt.args); got != tt.want {
				t.Errorf("FirstUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}
