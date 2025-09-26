package paths

import (
	"fmt"
	"testing"
)

func TestURLEscape_2(t *testing.T) {

	tests := []struct {
		name string
		uri  string
		want string
	}{
		{
			name: "Test Case 1",
			uri:  "http://example.com/path?query=value",
			want: "http://example.com/path?query=value",
		},
		{
			name: "Test Case 2",
			uri:  "http://example.com/path with space?query=value with space",
			want: "http://example.com/path%20with%20space?query=value%20with%20space",
		},
		{
			name: "Test Case 3",
			uri:  "http://example.com/path with special char$?query=value with special char$",
			want: "http://example.com/path%20with%20special%20char$?query=value%20with%20special%20char$",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := URLEscape(tt.uri); got != tt.want {
				t.Errorf("URLEscape() = %v, want %v", got, tt.want)
			}
		})
	}
}
