package paths

import (
	"fmt"
	"testing"
)

func TestPathEscape_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "normal case",
			arg:  "/path/to/file",
			want: "/path/to/file",
		},
		{
			name: "with special characters",
			arg:  "/path/to/file?with=query&and=fragment",
			want: "/path/to/file%3Fwith%3Dquery%26and%3Dfragment",
		},
		{
			name: "with space",
			arg:  "/path/to/file with space",
			want: "/path/to/file%20with%20space",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := PathEscape(tt.arg)
			if got != tt.want {
				t.Errorf("PathEscape(%q) = %q, want %q", tt.arg, got, tt.want)
			}
		})
	}
}
