package template

import (
	"fmt"
	"testing"
)

func TestCssEscaper_1(t *testing.T) {
	tests := []struct {
		name string
		args []any
		want string
	}{
		{
			name: "Test with single character",
			args: []any{'a'},
			want: "a",
		},
		{
			name: "Test with multiple characters",
			args: []any{'a', 'b', 'c'},
			want: "abc",
		},
		{
			name: "Test with special characters",
			args: []any{'<', '>', '&'},
			want: "&lt;&gt;&amp;",
		},
		{
			name: "Test with whitespace",
			args: []any{' ', '\n', '\t'},
			want: " &#xA;&#x9;",
		},
		{
			name: "Test with unicode characters",
			args: []any{'😀', '😃', '😄'},
			want: "&#x1F600;&#x1F603;&#x1F604;",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := cssEscaper(tt.args...); got != tt.want {
				t.Errorf("cssEscaper() = %v, want %v", got, tt.want)
			}
		})
	}
}
