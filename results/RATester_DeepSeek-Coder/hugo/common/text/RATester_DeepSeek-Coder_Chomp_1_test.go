package text

import (
	"fmt"
	"testing"
)

func TestChomp_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"Empty string", "", ""},
		{"No newline", "Hello, world", "Hello, world"},
		{"Trailing newline", "Hello, world\n", "Hello, world"},
		{"Trailing carriage return", "Hello, world\r", "Hello, world"},
		{"Trailing newline and carriage return", "Hello, world\n\r", "Hello, world"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Chomp(tt.arg); got != tt.want {
				t.Errorf("Chomp() = %v, want %v", got, tt.want)
			}
		})
	}
}
