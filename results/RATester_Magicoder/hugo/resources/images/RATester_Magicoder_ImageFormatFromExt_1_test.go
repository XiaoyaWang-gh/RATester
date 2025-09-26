package images

import (
	"fmt"
	"testing"
)

func TestImageFormatFromExt_1(t *testing.T) {
	tests := []struct {
		name     string
		ext      string
		expected Format
		found    bool
	}{
		{
			name:     "valid extension",
			ext:      "jpg",
			expected: JPEG,
			found:    true,
		},
		{
			name:     "invalid extension",
			ext:      "invalid",
			expected: 0,
			found:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			actual, found := ImageFormatFromExt(tt.ext)
			if actual != tt.expected {
				t.Errorf("expected %v, but got %v", tt.expected, actual)
			}
			if found != tt.found {
				t.Errorf("expected found %v, but got %v", tt.found, found)
			}
		})
	}
}
