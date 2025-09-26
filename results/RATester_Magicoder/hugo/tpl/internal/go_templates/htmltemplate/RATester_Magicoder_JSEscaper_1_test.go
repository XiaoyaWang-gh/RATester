package template

import (
	"fmt"
	"testing"
)

func TestJSEscaper_1(t *testing.T) {
	tests := []struct {
		name string
		args []any
		want string
	}{
		{
			name: "Test case 1",
			args: []any{"field name"},
			want: "field name",
		},
		{
			name: "Test case 2",
			args: []any{"field want"},
			want: "field want",
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
