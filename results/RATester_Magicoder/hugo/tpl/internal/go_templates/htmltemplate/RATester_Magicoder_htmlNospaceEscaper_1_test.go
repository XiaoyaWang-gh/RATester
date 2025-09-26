package template

import (
	"fmt"
	"testing"
)

func TesthtmlNospaceEscaper_1(t *testing.T) {
	tests := []struct {
		name string
		args []any
		want string
	}{
		{
			name: "Test case 1",
			args: []any{"<p>Hello, World!</p>"},
			want: "Hello,World!",
		},
		{
			name: "Test case 2",
			args: []any{"<p>Hello, <b>World</b>!</p>"},
			want: "Hello,World!",
		},
		{
			name: "Test case 3",
			args: []any{"<p>Hello, <b>World</b>!</p>", "text/html"},
			want: "Hello, World!",
		},
		{
			name: "Test case 4",
			args: []any{"<p>Hello, <b>World</b>!</p>", "text/plain"},
			want: "Hello, <b>World</b>!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := htmlNospaceEscaper(tt.args...); got != tt.want {
				t.Errorf("htmlNospaceEscaper() = %v, want %v", got, tt.want)
			}
		})
	}
}
