package hugo

import (
	"fmt"
	"testing"
)

func TestCompare_1(t *testing.T) {
	tests := []struct {
		name     string
		h        VersionString
		other    any
		expected int
	}{
		{
			name:     "Equal",
			h:        "1.2.3",
			other:    "1.2.3",
			expected: 0,
		},
		{
			name:     "Greater",
			h:        "1.2.4",
			other:    "1.2.3",
			expected: 1,
		},
		{
			name:     "Less",
			h:        "1.2.2",
			other:    "1.2.3",
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tt.h.Compare(tt.other)
			if result != tt.expected {
				t.Errorf("Expected %d, but got %d", tt.expected, result)
			}
		})
	}
}
