package tpl

import (
	"fmt"
	"testing"
)

func TestStripHTML_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "Test with simple string",
			arg:  "<p>Hello, World!</p>",
			want: "Hello, World!",
		},
		{
			name: "Test with complex string",
			arg:  "<p>Hello, <b>World</b>!</p>",
			want: "Hello, World!",
		},
		{
			name: "Test with no HTML",
			arg:  "Hello, World!",
			want: "Hello, World!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := StripHTML(tt.arg); got != tt.want {
				t.Errorf("StripHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}
