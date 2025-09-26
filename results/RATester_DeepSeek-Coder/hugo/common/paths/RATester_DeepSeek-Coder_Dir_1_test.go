package paths

import (
	"fmt"
	"testing"
)

func TestDir_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "Test with a valid path",
			in:   "/a/b/c/d.go",
			want: "/a/b/c",
		},
		{
			name: "Test with a valid path with trailing slash",
			in:   "/a/b/c/",
			want: "/a/b",
		},
		{
			name: "Test with a valid path with double trailing slash",
			in:   "/a/b/c//",
			want: "/a/b",
		},
		{
			name: "Test with a root path",
			in:   "/",
			want: "/",
		},
		{
			name: "Test with an empty path",
			in:   "",
			want: ".",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			got := pathBridge{}.Dir(tt.in)
			if got != tt.want {
				t.Errorf("pathBridge.Dir(%q) = %q; want %q", tt.in, got, tt.want)
			}
		})
	}
}
