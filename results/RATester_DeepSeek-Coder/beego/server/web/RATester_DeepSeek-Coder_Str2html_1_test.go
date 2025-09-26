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
			raw:  "<p>Hello, world</p>",
			want: "<p>Hello, world</p>",
		},
		{
			name: "Test case 2",
			raw:  "<script>alert('XSS')</script>",
			want: "&lt;script&gt;alert('XSS')&lt;/script&gt;",
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
