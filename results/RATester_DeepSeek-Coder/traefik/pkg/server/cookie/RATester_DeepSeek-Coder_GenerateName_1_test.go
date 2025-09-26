package cookie

import (
	"fmt"
	"testing"
)

func TestGenerateName_1(t *testing.T) {
	tests := []struct {
		name     string
		backend  string
		expected string
	}{
		{
			name:     "Test case 1",
			backend:  "backend1",
			expected: "_4397f7ee",
		},
		{
			name:     "Test case 2",
			backend:  "backend2",
			expected: "_4397f7ee",
		},
		{
			name:     "Test case 3",
			backend:  "backend3",
			expected: "_4397f7ee",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := GenerateName(test.backend)
			if result != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, result)
			}
		})
	}
}
