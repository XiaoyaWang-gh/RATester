package tplimpl

import (
	"fmt"
	"testing"
)

func TestNameIsText_1(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected string
		isText   bool
	}

	testCases := []testCase{
		{
			name:     "Text template",
			input:    "text-template",
			expected: "template",
			isText:   true,
		},
		{
			name:     "Non-text template",
			input:    "non-text-template",
			expected: "non-text-template",
			isText:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			handler := &templateHandler{}
			name, isText := handler.nameIsText(tc.input)
			if name != tc.expected {
				t.Errorf("Expected name %s, got %s", tc.expected, name)
			}
			if isText != tc.isText {
				t.Errorf("Expected isText %t, got %t", tc.isText, isText)
			}
		})
	}
}
