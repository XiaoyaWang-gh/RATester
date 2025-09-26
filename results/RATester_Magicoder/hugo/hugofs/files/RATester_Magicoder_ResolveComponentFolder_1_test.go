package files

import (
	"fmt"
	"testing"
)

func TestResolveComponentFolder_1(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     string
	}{
		{
			name:     "Test case 1",
			filename: "/path/to/component/file",
			want:     "/path/to/component/",
		},
		{
			name:     "Test case 2",
			filename: "/path/to/another/component/file",
			want:     "/path/to/another/component/",
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ResolveComponentFolder(tt.filename); got != tt.want {
				t.Errorf("ResolveComponentFolder() = %v, want %v", got, tt.want)
			}
		})
	}
}
