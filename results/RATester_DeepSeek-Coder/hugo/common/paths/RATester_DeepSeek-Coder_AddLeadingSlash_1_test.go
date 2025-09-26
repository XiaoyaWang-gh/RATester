package paths

import (
	"fmt"
	"testing"
)

func TestAddLeadingSlash_1(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "Test with leading slash",
			path: "/test",
			want: "/test",
		},
		{
			name: "Test without leading slash",
			path: "test",
			want: "/test",
		},
		{
			name: "Test with empty string",
			path: "",
			want: "/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := AddLeadingSlash(tt.path); got != tt.want {
				t.Errorf("AddLeadingSlash() = %v, want %v", got, tt.want)
			}
		})
	}
}
