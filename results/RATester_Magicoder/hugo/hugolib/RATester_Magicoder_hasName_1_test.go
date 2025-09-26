package hugolib

import (
	"fmt"
	"testing"
)

func TesthasName_1(t *testing.T) {
	handler := &shortcodeHandler{
		nameSet: map[string]bool{
			"test": true,
		},
	}

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Existing name",
			input:    "test",
			expected: true,
		},
		{
			name:     "Non-existing name",
			input:    "nonexistent",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := handler.hasName(test.input)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
