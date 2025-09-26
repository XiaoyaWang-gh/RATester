package text

import (
	"fmt"
	"testing"
)

func TestPuts_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"empty string", "", ""},
		{"string ends with newline", "Hello\n", "Hello\n"},
		{"string does not end with newline", "Hello", "Hello\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Puts(tt.arg); got != tt.want {
				t.Errorf("Puts() = %v, want %v", got, tt.want)
			}
		})
	}
}
