package web

import (
	"fmt"
	"html/template"
	"testing"
)

func TestStr2html_1(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want template.HTML
	}{
		{
			name: "Test case 1",
			raw:  "<script>alert('Hello, world!')</script>",
			want: "&lt;script&gt;alert('Hello, world!')&lt;/script&gt;",
		},
		{
			name: "Test case 2",
			raw:  "<b>Hello, world!</b>",
			want: "<b>Hello, world!</b>",
		},
		{
			name: "Test case 3",
			raw:  "Hello, world!",
			want: "Hello, world!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Str2html(tt.raw); got != tt.want {
				t.Errorf("Str2html() = %v, want %v", got, tt.want)
			}
		})
	}
}
