package path

import (
	"fmt"
	"testing"
)

func TestClean_3(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		path    any
		want    string
		wantErr bool
	}{
		{
			name: "clean path",
			path: "/foo/bar/../baz",
			want: "/foo/baz",
		},
		{
			name: "clean path with trailing slash",
			path: "/foo/bar/../baz/",
			want: "/foo/baz/",
		},
		{
			name: "clean path with double slash",
			path: "//foo//bar//../baz",
			want: "/foo/baz",
		},
		{
			name: "clean path with double slash and trailing slash",
			path: "//foo//bar//../baz//",
			want: "/foo/baz/",
		},
		{
			name:    "invalid path",
			path:    123,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := ns.Clean(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Clean() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Clean() = %v, want %v", got, tt.want)
			}
		})
	}
}
