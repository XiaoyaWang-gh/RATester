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
			args: []any{"field name"},
			want: "field%20name",
		},
		{
			name: "Test case 2",
			args: []any{"field want"},
			want: "field%20want",
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
