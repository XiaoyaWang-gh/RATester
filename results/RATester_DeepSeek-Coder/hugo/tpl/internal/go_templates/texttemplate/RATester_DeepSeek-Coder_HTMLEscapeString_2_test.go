package template

import (
	"fmt"
	"testing"
)

func TestHTMLEscapeString_2(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "Test with simple string",
			arg:  "Hello, World",
			want: "Hello, World",
		},
		{
			name: "Test with special characters",
			arg:  "<script>alert('XSS')</script>",
			want: "&lt;script&gt;alert('XSS')&lt;/script&gt;",
		},
		{
			name: "Test with null character",
			arg:  "\x00",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := HTMLEscapeString(tt.arg); got != tt.want {
				t.Errorf("HTMLEscapeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
