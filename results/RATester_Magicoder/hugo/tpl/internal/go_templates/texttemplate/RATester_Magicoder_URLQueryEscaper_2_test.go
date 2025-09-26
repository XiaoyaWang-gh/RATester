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
			args: []any{"field name", "value"},
			want: "field+name=value",
		},
		{
			name: "Test case 2",
			args: []any{"field2", "value2"},
			want: "field2=value2",
		},
		// Add more test cases as needed
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
