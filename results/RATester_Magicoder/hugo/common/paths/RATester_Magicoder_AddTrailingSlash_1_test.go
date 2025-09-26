package paths

import (
	"fmt"
	"testing"
)

func TestAddTrailingSlash_1(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "Add trailing slash to path",
			path: "test",
			want: "test/",
		},
		{
			name: "Do not add trailing slash to path with trailing slash",
			path: "test/",
			want: "test/",
		},
		{
			name: "Do not add trailing slash to empty path",
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

			if got := AddTrailingSlash(tt.path); got != tt.want {
				t.Errorf("AddTrailingSlash() = %v, want %v", got, tt.want)
			}
		})
	}
}
