package images

import (
	"fmt"
	"testing"
)

func TestImageFormatFromMediaSubType_1(t *testing.T) {
	tests := []struct {
		name     string
		sub      string
		expected Format
		found    bool
	}{
		{
			name:     "valid subtype",
			sub:      "jpeg",
			expected: JPEG,
			found:    true,
		},
		{
			name:     "invalid subtype",
			sub:      "invalid",
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

			actual, found := ImageFormatFromMediaSubType(tt.sub)
			if actual != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
			if found != tt.found {
				t.Errorf("expected found %v, got %v", tt.found, found)
			}
		})
	}
}
