package template

import (
	"fmt"
	"testing"
)

func TesturlProcessor_1(t *testing.T) {
	tests := []struct {
		name string
		norm bool
		args []any
		want string
	}{
		{
			name: "Test Case 1",
			norm: true,
			args: []any{"http://example.com"},
			want: "http://example.com",
		},
		{
			name: "Test Case 2",
			norm: false,
			args: []any{"http://example.com"},
			want: "http://example.com",
		},
		{
			name: "Test Case 3",
			norm: true,
			args: []any{"http://example.com", contentTypeURL},
			want: "http://example.com",
		},
		{
			name: "Test Case 4",
			norm: false,
			args: []any{"http://example.com", contentTypeURL},
			want: "http://example.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := urlProcessor(tt.norm, tt.args...); got != tt.want {
				t.Errorf("urlProcessor() = %v, want %v", got, tt.want)
			}
		})
	}
}
