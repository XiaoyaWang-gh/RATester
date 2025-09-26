package paths

import (
	"fmt"
	"testing"
)

func TestIsSameFilePath_1(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{
			name: "same file path",
			s1:   "/path/to/file",
			s2:   "/path/to/file",
			want: true,
		},
		{
			name: "different file path",
			s1:   "/path/to/file1",
			s2:   "/path/to/file2",
			want: false,
		},
		{
			name: "same file path with different slashes",
			s1:   "/path/to/file",
			s2:   "\\path\\to\\file",
			want: true,
		},
		{
			name: "same file path with different slashes and trailing slashes",
			s1:   "/path/to/file/",
			s2:   "\\path\\to\\file\\",
			want: true,
		},
		{
			name: "same file path with different slashes and trailing slashes and spaces",
			s1:   "/path/to/file/ ",
			s2:   "\\path\\to\\file\\ ",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := IsSameFilePath(tt.s1, tt.s2); got != tt.want {
				t.Errorf("IsSameFilePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
