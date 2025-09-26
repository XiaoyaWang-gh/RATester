package paths

import (
	"fmt"
	"testing"
)

func TestToSlashTrimTrailing_1(t *testing.T) {
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
			arg:  "/path/to/file/",
			want: "/path/to/file",
		},
		{
			name: "Test case 3",
			arg:  "path/to/file",
			want: "path/to/file",
		},
		{
			name: "Test case 4",
			arg:  "path/to/file/",
			want: "path/to/file",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ToSlashTrimTrailing(tt.arg); got != tt.want {
				t.Errorf("ToSlashTrimTrailing() = %v, want %v", got, tt.want)
			}
		})
	}
}
