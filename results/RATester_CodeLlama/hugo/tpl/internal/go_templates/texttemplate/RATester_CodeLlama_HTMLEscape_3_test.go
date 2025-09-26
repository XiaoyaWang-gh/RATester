package template

import (
	"bytes"
	"fmt"
	"testing"
)

func TestHTMLEscape_3(t *testing.T) {
	var tests = []struct {
		name string
		b    []byte
		want string
	}{
		{
			name: "empty",
			b:    []byte(""),
			want: "",
		},
		{
			name: "no escape",
			b:    []byte("hello"),
			want: "hello",
		},
		{
			name: "escape",
			b:    []byte("hello\"world'&<>`"),
			want: "hello&quot;world&#39;&amp;&lt;&gt;`",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			var buf bytes.Buffer
			HTMLEscape(&buf, tt.b)
			if got := buf.String(); got != tt.want {
				t.Errorf("HTMLEscape() = %q, want %q", got, tt.want)
			}
		})
	}
}
