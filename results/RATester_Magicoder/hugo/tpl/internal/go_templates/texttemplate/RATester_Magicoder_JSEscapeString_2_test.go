package template

import (
	"fmt"
	"testing"
)

func TestJSEscapeString_2(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"empty string", "", ""},
		{"no special characters", "hello", "hello"},
		{"special characters", "hello<>&\"'", "hello\\u003C\\u003E\\u0026\\u0022\\u0027"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := JSEscapeString(tt.arg); got != tt.want {
				t.Errorf("JSEscapeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
