package template

import (
	"fmt"
	"testing"
)

func TesturlEscaper_1(t *testing.T) {
	tests := []struct {
		name string
		args []any
		want string
	}{
		{
			name: "Test Case 1",
			args: []any{"field name"},
			want: "field%20name",
		},
		{
			name: "Test Case 2",
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

			if got := urlEscaper(tt.args...); got != tt.want {
				t.Errorf("urlEscaper() = %v, want %v", got, tt.want)
			}
		})
	}
}
