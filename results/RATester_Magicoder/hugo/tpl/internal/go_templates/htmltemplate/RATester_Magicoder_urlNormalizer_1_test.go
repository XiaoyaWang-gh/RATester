package template

import (
	"fmt"
	"testing"
)

func TesturlNormalizer_1(t *testing.T) {
	tests := []struct {
		name string
		args []any
		want string
	}{
		{
			name: "Test Case 1",
			args: []any{"https://example.com", "field name"},
			want: "https://example.com/field name",
		},
		{
			name: "Test Case 2",
			args: []any{"https://example.com/", "field name"},
			want: "https://example.com/field name",
		},
		{
			name: "Test Case 3",
			args: []any{"https://example.com", "/field name"},
			want: "https://example.com/field name",
		},
		{
			name: "Test Case 4",
			args: []any{"https://example.com/", "/field name"},
			want: "https://example.com/field name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := urlNormalizer(tt.args...); got != tt.want {
				t.Errorf("urlNormalizer() = %v, want %v", got, tt.want)
			}
		})
	}
}
