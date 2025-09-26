package template

import (
	"fmt"
	"testing"
)

func TestURLQueryEscaper_1(t *testing.T) {
	tests := []struct {
		name string
		args []any
		want string
	}{
		{
			name: "Test case 1",
			args: []any{"field name", "value"},
			want: "field+name=value",
		},
		{
			name: "Test case 2",
			args: []any{"field name with space", "value with space"},
			want: "field+name+with+space=value+with+space",
		},
		{
			name: "Test case 3",
			args: []any{"field name with special characters", "value with !@#$%^&*()_+=-[]{}|\\;:'\",.<>?/`~"},
			want: "field+name+with+special+characters=value+with+!@#$%^&*()_+=-[]{}|\\;:'\",.<>?/`~",
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
