package files

import (
	"fmt"
	"testing"
)

func TestIsComponentFolder_1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Test case 1: Existing component folder",
			input:    "existing_component_folder",
			expected: true,
		},
		{
			name:     "Test case 2: Non-existing component folder",
			input:    "non_existing_component_folder",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := IsComponentFolder(tt.input)
			if got != tt.expected {
				t.Errorf("IsComponentFolder(%s) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
