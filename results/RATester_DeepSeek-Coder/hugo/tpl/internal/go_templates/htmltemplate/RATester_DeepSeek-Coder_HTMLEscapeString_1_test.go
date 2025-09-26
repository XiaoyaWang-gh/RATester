package template

import (
	"fmt"
	"testing"
)

func TestHTMLEscapeString_1(t *testing.T) {
	testCases := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "Test with simple string",
			arg:  "<script>alert('XSS');</script>",
			want: "&lt;script&gt;alert('XSS');&lt;/script&gt;",
		},
		{
			name: "Test with special characters",
			arg:  "&<>\"'",
			want: "&amp;&lt;&gt;&quot;&#39;",
		},
		{
			name: "Test with empty string",
			arg:  "",
			want: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := HTMLEscapeString(tc.arg)
			if got != tc.want {
				t.Errorf("HTMLEscapeString(%q) = %q; want %q", tc.arg, got, tc.want)
			}
		})
	}
}
