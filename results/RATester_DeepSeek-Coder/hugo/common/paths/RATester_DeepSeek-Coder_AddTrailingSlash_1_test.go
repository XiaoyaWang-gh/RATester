package paths

import (
	"fmt"
	"testing"
)

func TestAddTrailingSlash_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty string", "", "/"},
		{"String without trailing slash", "test", "test/"},
		{"String with trailing slash", "test/", "test/"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := AddTrailingSlash(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
