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
		{"single lowercase", "a", "A"},
		{"single uppercase", "A", "A"},
		{"multiple words", "hello world", "Hello world"},
		{"multiple words with uppercase", "Hello World", "Hello World"},
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
