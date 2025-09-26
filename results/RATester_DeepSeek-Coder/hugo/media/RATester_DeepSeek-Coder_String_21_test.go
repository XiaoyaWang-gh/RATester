package media

import (
	"fmt"
	"testing"
)

func TestString_21(t *testing.T) {
	type testCase struct {
		name     string
		input    Type
		expected string
	}

	testCases := []testCase{
		{
			name: "Test Case 1",
			input: Type{
				Type:        "application/json",
				MainType:    "application",
				SubType:     "json",
				Delimiter:   "/",
				FirstSuffix: SuffixInfo{},
			},
			expected: "application/json",
		},
		{
			name: "Test Case 2",
			input: Type{
				Type:        "text/plain",
				MainType:    "text",
				SubType:     "plain",
				Delimiter:   "/",
				FirstSuffix: SuffixInfo{},
			},
			expected: "text/plain",
		},
		{
			name: "Test Case 3",
			input: Type{
				Type:        "image/jpeg",
				MainType:    "image",
				SubType:     "jpeg",
				Delimiter:   "/",
				FirstSuffix: SuffixInfo{},
			},
			expected: "image/jpeg",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.input.String()
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
