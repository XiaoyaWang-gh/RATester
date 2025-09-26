package hugo

import (
	"fmt"
	"testing"
)

func TestEq_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		h        VersionString
		other    any
		expected bool
	}{
		{
			name:     "Equal strings",
			h:        "1.2.3",
			other:    "1.2.3",
			expected: true,
		},
		{
			name:     "Different strings",
			h:        "1.2.3",
			other:    "1.2.4",
			expected: false,
		},
		{
			name:     "Different types",
			h:        "1.2.3",
			other:    123,
			expected: false,
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

			result := tt.h.Eq(tt.other)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
