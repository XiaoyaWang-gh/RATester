package template

import (
	"fmt"
	"testing"
)

func TestJSEscaper_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args []any
		want string
	}{
		{
			name: "Test case 1",
			args: []any{"<script>alert('XSS')</script>"},
			want: "\\u003Cscript\\u003Ealert('XSS')\\u003C/script\\u003E",
		},
		{
			name: "Test case 2",
			args: []any{">", "<"},
			want: "\\u003E \\u003C",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := JSEscaper(tt.args...); got != tt.want {
				t.Errorf("JSEscaper() = %v, want %v", got, tt.want)
			}
		})
	}
}
