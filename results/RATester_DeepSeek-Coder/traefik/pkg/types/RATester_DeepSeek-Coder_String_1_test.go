package types

import (
	"fmt"
	"testing"
)

func TestString_1(t *testing.T) {
	tests := []struct {
		name string
		f    FileOrContent
		want string
	}{
		{
			name: "Test with file path",
			f:    FileOrContent("/path/to/file"),
			want: "/path/to/file",
		},
		{
			name: "Test with file content",
			f:    FileOrContent("file content"),
			want: "file content",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.f.String(); got != tt.want {
				t.Errorf("FileOrContent.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
