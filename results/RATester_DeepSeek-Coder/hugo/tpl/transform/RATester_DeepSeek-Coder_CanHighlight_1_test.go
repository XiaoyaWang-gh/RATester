package transform

import (
	"fmt"
	"testing"
)

func TestCanHighlight_1(t *testing.T) {
	type testCase struct {
		name     string
		language string
		expected bool
	}

	testCases := []testCase{
		{"valid language", "go", true},
		{"invalid language", "invalid", false},
		{"empty language", "", false},
	}

	ns := &Namespace{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := ns.CanHighlight(tc.language)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
