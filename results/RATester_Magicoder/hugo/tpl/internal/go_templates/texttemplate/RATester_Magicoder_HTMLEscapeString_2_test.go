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
			name: "Test case 1",
			arg:  "<script>alert('Hello, world!')</script>",
			want: "&lt;script&gt;alert('Hello, world!')&lt;/script&gt;",
		},
		{
			name: "Test case 2",
			arg:  "Hello, world!",
			want: "Hello, world!",
		},
		{
			name: "Test case 3",
			arg:  "<script>alert('Hello, world!')</script>",
			want: "&lt;script&gt;alert('Hello, world!')&lt;/script&gt;",
		},
		{
			name: "Test case 4",
			arg:  "Hello, world!",
			want: "Hello, world!",
		},
		{
			name: "Test case 5",
			arg:  "<script>alert('Hello, world!')</script>",
			want: "&lt;script&gt;alert('Hello, world!')&lt;/script&gt;",
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
