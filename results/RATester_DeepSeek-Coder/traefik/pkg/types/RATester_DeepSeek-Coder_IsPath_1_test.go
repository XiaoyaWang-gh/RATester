package types

import (
	"fmt"
	"testing"
)

func TestIsPath_1(t *testing.T) {
	tests := []struct {
		name string
		f    FileOrContent
		want bool
	}{
		{
			name: "Test with existing file path",
			f:    FileOrContent("/etc/passwd"),
			want: true,
		},
		{
			name: "Test with non-existing file path",
			f:    FileOrContent("/non-existing-file"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.f.IsPath(); got != tt.want {
				t.Errorf("FileOrContent.IsPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
