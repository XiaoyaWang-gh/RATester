package template

import (
	"fmt"
	"testing"
)

func TestcssEscaper_1(t *testing.T) {
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
			args: []any{"field name", "field want"},
			want: "field name field want",
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

			if got := cssEscaper(tt.args...); got != tt.want {
				t.Errorf("cssEscaper() = %v, want %v", got, tt.want)
			}
		})
	}
}
