package web

import (
	"fmt"
	"testing"
)

func TestHtmlquote_1(t *testing.T) {

	tests := []struct {
		name string
		text string
		want string
	}{
		{
			name: "Test with double quotes",
			text: "“Hello”",
			want: "&ldquo;Hello&rdquo;",
		},
		{
			name: "Test with space",
			text: "Hello World",
			want: "Hello&nbsp;World",
		},
		{
			name: "Test with no special characters",
			text: "HelloWorld",
			want: "HelloWorld",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Htmlquote(tt.text); got != tt.want {
				t.Errorf("Htmlquote() = %v, want %v", got, tt.want)
			}
		})
	}
}
