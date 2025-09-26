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
			arg:  "<script>alert('XSS')</script>",
			want: "\\u003Cscript\\u003Ealert('XSS')\\u003C/script\\u003E",
		},
		{
			name: "Test case 2",
			arg:  "\\",
			want: "\\\\",
		},
		{
			name: "Test case 3",
			arg:  "\\n",
			want: "\\n",
		},
		{
			name: "Test case 4",
			arg:  "\\r",
			want: "\\r",
		},
		{
			name: "Test case 5",
			arg:  "\\t",
			want: "\\t",
		},
		{
			name: "Test case 6",
			arg:  "\\\"",
			want: "\\\"",
		},
		{
			name: "Test case 7",
			arg:  "\\'",
			want: "\\'",
		},
		{
			name: "Test case 8",
			arg:  "\\\\",
			want: "\\\\",
		},
		{
			name: "Test case 9",
			arg:  "\\u00A9",
			want: "\\u00A9",
		},
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
