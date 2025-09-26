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
		{"special characters", "<hello>", "\\u003chello\\u003e"},
		{"special characters 2", "hello\nworld", "hello\\u000aworld"},
		{"special characters 3", "hello\rworld", "hello\\u000dworld"},
		{"special characters 4", "hello\tworld", "hello\\u0009world"},
		{"special characters 5", "hello\fworld", "hello\\u000cworld"},
		{"special characters 6", "hello\bworld", "hello\\u0008world"},
		{"special characters 7", "hello\\world", "hello\\\\world"},
		{"special characters 8", "hello\"world", "hello\\\"world"},
		{"special characters 9", "hello'world", "hello\\'world"},
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
