package helpers

import (
	"fmt"
	"testing"
)

func TestaddTrailingFileSeparator_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "Test with trailing slash",
			arg:  "/path/to/dir/",
			want: "/path/to/dir/",
		},
		{
			name: "Test without trailing slash",
			arg:  "/path/to/dir",
			want: "/path/to/dir/",
		},
		{
			name: "Test with empty string",
			arg:  "",
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

			if got := addTrailingFileSeparator(tt.arg); got != tt.want {
				t.Errorf("addTrailingFileSeparator() = %v, want %v", got, tt.want)
			}
		})
	}
}
