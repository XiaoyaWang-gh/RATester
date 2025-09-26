package constants

import (
	"fmt"
	"testing"
)

func TestIsFieldRelOrPermalink_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Test Case 1", "relpermalink", true},
		{"Test Case 2", "permalink", true},
		{"Test Case 3", "other", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := IsFieldRelOrPermalink(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
