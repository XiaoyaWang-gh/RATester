package terminal

import (
	"fmt"
	"testing"
)

func TestdoublePercent_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testCases := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"Hello%World", "Hello%%World"},
		{"%%", "%%%%"},
		{"100%%", "100%%"},
		{"%%100", "%%100"},
	}

	for _, tc := range testCases {
		result := doublePercent(tc.input)
		if result != tc.expected {
			t.Errorf("Expected %s, but got %s", tc.expected, result)
		}
	}
}
