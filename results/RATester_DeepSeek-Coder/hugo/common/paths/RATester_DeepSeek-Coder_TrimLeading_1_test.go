package paths

import (
	"fmt"
	"testing"
)

func TestTrimLeading_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"no leading slash", "test", "test"},
		{"leading slash", "/test", "test"},
		{"multiple leading slashes", "////test", "test"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := TrimLeading(tc.input)
			if result != tc.expected {
				t.Errorf("Expected '%s', got '%s'", tc.expected, result)
			}
		})
	}
}
