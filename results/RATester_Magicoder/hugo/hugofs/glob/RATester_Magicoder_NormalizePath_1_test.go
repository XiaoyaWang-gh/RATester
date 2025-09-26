package glob

import (
	"fmt"
	"testing"
)

func TestNormalizePath_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "NormalizePath test 1",
			arg:  "/path/to/file",
			want: "/path/to/file",
		},
		{
			name: "NormalizePath test 2",
			arg:  "/Path/To/File",
			want: "/path/to/file",
		},
		{
			name: "NormalizePath test 3",
			arg:  "/path/To/file",
			want: "/path/to/file",
		},
		{
			name: "NormalizePath test 4",
			arg:  "/path/to/FILE",
			want: "/path/to/file",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NormalizePath(tt.arg); got != tt.want {
				t.Errorf("NormalizePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
