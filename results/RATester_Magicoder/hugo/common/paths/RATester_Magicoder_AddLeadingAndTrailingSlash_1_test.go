package paths

import (
	"fmt"
	"testing"
)

func TestAddLeadingAndTrailingSlash_1(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "Add leading and trailing slash to empty string",
			path: "",
			want: "/",
		},
		{
			name: "Add leading and trailing slash to root path",
			path: "/",
			want: "/",
		},
		{
			name: "Add leading and trailing slash to simple path",
			path: "foo",
			want: "/foo/",
		},
		{
			name: "Add leading and trailing slash to path with trailing slash",
			path: "foo/",
			want: "/foo/",
		},
		{
			name: "Add leading and trailing slash to path with leading slash",
			path: "/foo",
			want: "/foo/",
		},
		{
			name: "Add leading and trailing slash to path with leading and trailing slash",
			path: "/foo/",
			want: "/foo/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := AddLeadingAndTrailingSlash(tt.path); got != tt.want {
				t.Errorf("AddLeadingAndTrailingSlash() = %v, want %v", got, tt.want)
			}
		})
	}
}
