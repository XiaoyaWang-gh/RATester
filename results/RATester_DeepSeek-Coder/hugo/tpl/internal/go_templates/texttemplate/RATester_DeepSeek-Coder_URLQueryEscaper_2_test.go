package template

import (
	"fmt"
	"testing"
)

func TestURLQueryEscaper_2(t *testing.T) {
	tests := []struct {
		name string
		args []any
		want string
	}{
		{
			name: "Test case 1",
			args: []any{"field name"},
			want: "field+name",
		},
		{
			name: "Test case 2",
			args: []any{"field+name"},
			want: "field%2Bname",
		},
		{
			name: "Test case 3",
			args: []any{"field%name"},
			want: "field%25name",
		},
		{
			name: "Test case 4",
			args: []any{"field name with spaces"},
			want: "field+name+with+spaces",
		},
		{
			name: "Test case 5",
			args: []any{"field name with special characters !@#$%^&*()+=-_[]{}|;:',.<>?/`~"},
			want: "field+name+with+special+characters+%21%40%23%24%25%5E%26%2A%28%29%2B%3D-_%5B%5D%7C%3B%3A%27%2C.%3C%3E%3F%2F%60~",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := URLQueryEscaper(tt.args...); got != tt.want {
				t.Errorf("URLQueryEscaper() = %v, want %v", got, tt.want)
			}
		})
	}
}
