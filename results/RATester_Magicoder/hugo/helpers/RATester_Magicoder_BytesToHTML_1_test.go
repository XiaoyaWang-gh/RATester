package helpers

import (
	"fmt"
	"html/template"
	"testing"
)

func TestBytesToHTML_1(t *testing.T) {
	tests := []struct {
		name string
		b    []byte
		want template.HTML
	}{
		{
			name: "Test Case 1",
			b:    []byte("<p>Hello, World!</p>"),
			want: "<p>Hello, World!</p>",
		},
		{
			name: "Test Case 2",
			b:    []byte("<script>alert('Hello, World!');</script>"),
			want: "<script>alert('Hello, World!');</script>",
		},
		{
			name: "Test Case 3",
			b:    []byte("<p>Hello, </p>\n<p>World!</p>"),
			want: "<p>Hello, </p>\n<p>World!</p>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := BytesToHTML(tt.b); got != tt.want {
				t.Errorf("BytesToHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}
