package parse

import (
	"fmt"
	"testing"
)

func TestHasRightTrimMarker_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", false},
		{"single character", "a", false},
		{"two characters, no trim marker", " a", false},
		{"two characters, with trim marker", " a ", true},
		{"three characters, with trim marker", " aa ", true},
		{"four characters, with trim marker", " aaa ", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := hasRightTrimMarker(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
