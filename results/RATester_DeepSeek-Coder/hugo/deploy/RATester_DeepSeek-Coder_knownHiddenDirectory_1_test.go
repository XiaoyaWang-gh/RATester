package deploy

import (
	"fmt"
	"testing"
)

func TestKnownHiddenDirectory_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Test Case 1", ".well-known", true},
		{"Test Case 2", "not-known", false},
		{"Test Case 3", "", false},
		{"Test Case 4", " ", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := knownHiddenDirectory(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
