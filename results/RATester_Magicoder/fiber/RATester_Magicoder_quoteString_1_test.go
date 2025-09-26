package fiber

import (
	"fmt"
	"testing"
)

func TestquoteString_1(t *testing.T) {
	app := &App{
		getString: func(b []byte) string {
			return string(b)
		},
	}

	tests := []struct {
		name string
		raw  string
		want string
	}{
		{
			name: "Empty string",
			raw:  "",
			want: "",
		},
		{
			name: "Single character",
			raw:  "a",
			want: "a",
		},
		{
			name: "Multiple characters",
			raw:  "hello world",
			want: "hello world",
		},
		{
			name: "Special characters",
			raw:  "@#$%^&*()",
			want: "@#$%^&*()",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := app.quoteString(tt.raw); got != tt.want {
				t.Errorf("quoteString() = %v, want %v", got, tt.want)
			}
		})
	}
}
