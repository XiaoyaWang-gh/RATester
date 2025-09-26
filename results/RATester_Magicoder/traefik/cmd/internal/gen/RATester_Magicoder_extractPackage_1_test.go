package main

import (
	"fmt"
	"go/types"
	"testing"
)

func TestExtractPackage_1(t *testing.T) {
	type testCase struct {
		name     string
		input    types.Type
		expected string
	}

	testCases := []testCase{
		{
			name:     "Named",
			input:    &types.Named{},
			expected: "path/to/package",
		},
		{
			name:     "Slice",
			input:    &types.Slice{},
			expected: "path/to/package",
		},
		{
			name:     "Map",
			input:    &types.Map{},
			expected: "path/to/package",
		},
		{
			name:     "Pointer",
			input:    &types.Pointer{},
			expected: "path/to/package",
		},
		{
			name:     "Default",
			input:    &types.Basic{},
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := extractPackage(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
