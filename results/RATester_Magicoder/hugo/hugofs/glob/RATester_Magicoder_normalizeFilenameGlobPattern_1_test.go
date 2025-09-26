package glob

import (
	"fmt"
	"testing"
)

func TestnormalizeFilenameGlobPattern_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testCases := []struct {
		input    string
		expected string
	}{
		{"", "/"},
		{"/", "/"},
		{"test", "/test"},
		{"/test", "/test"},
		{"test/", "/test/"},
		{"/test/", "/test/"},
	}

	for _, tc := range testCases {
		result := normalizeFilenameGlobPattern(tc.input)
		if result != tc.expected {
			t.Errorf("Expected %s, got %s", tc.expected, result)
		}
	}
}
