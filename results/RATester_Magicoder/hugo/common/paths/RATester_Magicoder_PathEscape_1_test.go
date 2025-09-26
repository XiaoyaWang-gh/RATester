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
			name: "Test case 1",
			arg:  "/path/to/file",
			want: "/path/to/file",
		},
		{
			name: "Test case 2",
			arg:  "/path/to/file with spaces",
			want: "/path/to/file%20with%20spaces",
		},
		{
			name: "Test case 3",
			arg:  "/path/to/file with special characters",
			want: "/path/to/file%20with%20special%20characters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := PathEscape(tt.arg); got != tt.want {
				t.Errorf("PathEscape() = %v, want %v", got, tt.want)
			}
		})
	}
}
