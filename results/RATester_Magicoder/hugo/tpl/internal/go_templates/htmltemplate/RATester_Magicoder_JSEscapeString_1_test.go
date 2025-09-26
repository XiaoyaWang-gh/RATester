package template

import (
	"fmt"
	"testing"
)

func TestJSEscapeString_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "Test case 1",
			arg:  "<script>alert('Hello, world!')</script>",
			want: "\\u003Cscript\\u003Ealert('Hello, world!')\\u003C/script\\u003E",
		},
		{
			name: "Test case 2",
			arg:  "<",
			want: "\\u003C",
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

			if got := JSEscapeString(tt.arg); got != tt.want {
				t.Errorf("JSEscapeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
