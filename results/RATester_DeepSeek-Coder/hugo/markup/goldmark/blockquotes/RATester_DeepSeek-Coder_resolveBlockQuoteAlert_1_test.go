package blockquotes

import (
	"fmt"
	"testing"
)

func TestResolveBlockQuoteAlert_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected blockQuoteAlert
	}{
		{
			name:  "valid block quote alert",
			input: "> [info] + This is a valid block quote alert",
			expected: blockQuoteAlert{
				typ:   "info",
				sign:  "+",
				title: "This is a valid block quote alert",
			},
		},
		{
			name:  "valid block quote alert with no title",
			input: "> [info] -",
			expected: blockQuoteAlert{
				typ:   "info",
				sign:  "-",
				title: "",
			},
		},
		{
			name:     "invalid block quote alert",
			input:    "> [info This is an invalid block quote alert",
			expected: blockQuoteAlert{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := resolveBlockQuoteAlert(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
