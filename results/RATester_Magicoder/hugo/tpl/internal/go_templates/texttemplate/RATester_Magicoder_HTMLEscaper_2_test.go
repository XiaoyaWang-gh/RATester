package template

import (
	"fmt"
	"testing"
)

func TestHTMLEscaper_2(t *testing.T) {
	tests := []struct {
		name string
		args []any
		want string
	}{
		{
			name: "Test case 1",
			args: []any{"<script>alert('XSS')</script>"},
			want: "&lt;script&gt;alert('XSS')&lt;/script&gt;",
		},
		{
			name: "Test case 2",
			args: []any{"<b>bold</b>"},
			want: "&lt;b&gt;bold&lt;/b&gt;",
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := HTMLEscaper(tt.args...); got != tt.want {
				t.Errorf("HTMLEscaper() = %v, want %v", got, tt.want)
			}
		})
	}
}
