package web

import (
	"fmt"
	"testing"
)

func TestIsValidForInput_1(t *testing.T) {
	testCases := []struct {
		name     string
		fType    string
		expected bool
	}{
		{"Test Case 1", "text", true},
		{"Test Case 2", "password", true},
		{"Test Case 3", "checkbox", true},
		{"Test Case 4", "radio", true},
		{"Test Case 5", "submit", true},
		{"Test Case 6", "reset", true},
		{"Test Case 7", "hidden", true},
		{"Test Case 8", "image", true},
		{"Test Case 9", "file", true},
		{"Test Case 10", "button", true},
		{"Test Case 11", "search", true},
		{"Test Case 12", "email", true},
		{"Test Case 13", "url", true},
		{"Test Case 14", "tel", true},
		{"Test Case 15", "number", true},
		{"Test Case 16", "range", true},
		{"Test Case 17", "date", true},
		{"Test Case 18", "month", true},
		{"Test Case 19", "week", true},
		{"Test Case 20", "time", true},
		{"Test Case 21", "datetime", true},
		{"Test Case 22", "datetime-local", true},
		{"Test Case 23", "color", true},
		{"Test Case 24", "invalid", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := isValidForInput(tc.fType)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
