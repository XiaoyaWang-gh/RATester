package paths

import (
	"fmt"
	"testing"
)

func TestTrimTrailing_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty string", "", ""},
		{"No trailing slash", "example.com", "example.com"},
		{"Trailing slash", "example.com/", "example.com"},
		{"Multiple trailing slashes", "example.com///", "example.com"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := TrimTrailing(tc.input)
			if result != tc.expected {
				t.Errorf("Expected '%s', but got '%s'", tc.expected, result)
			}
		})
	}
}
