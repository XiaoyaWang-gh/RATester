package context

import (
	"fmt"
	"testing"
)

func TestSanitizeValue_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty string", "", ""},
		{"No special characters", "test", "test"},
		{"Special characters", "test@#$%^&*()_+`~-=[]\\\";',./{}|\":<>?", "test"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := sanitizeValue(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
