package images

import (
	"fmt"
	"testing"
)

func TestImageFormatFromMediaSubType_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		sub      string
		expected Format
		found    bool
	}{
		{
			name:     "JPEG",
			sub:      "jpeg",
			expected: JPEG,
			found:    true,
		},
		{
			name:     "PNG",
			sub:      "png",
			expected: PNG,
			found:    true,
		},
		{
			name:     "Unknown",
			sub:      "unknown",
			expected: 0,
			found:    false,
		},
	}

	for _, tt := range tests {
		tt := tt // NOTE: https://golang.org/doc/faq#closures_and_goroutines
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			actual, found := ImageFormatFromMediaSubType(tt.sub)
			if actual != tt.expected || found != tt.found {
				t.Errorf("expected (%v, %v), got (%v, %v)", tt.expected, tt.found, actual, found)
			}
		})
	}
}
