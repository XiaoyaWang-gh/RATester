package pageparser

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/parser/metadecoders"
)

func TestFormatFromFrontMatterType_1(t *testing.T) {
	type testCase struct {
		name     string
		input    ItemType
		expected metadecoders.Format
	}

	testCases := []testCase{
		{
			name:     "Test Case 1: TypeFrontMatterJSON",
			input:    TypeFrontMatterJSON,
			expected: metadecoders.JSON,
		},
		{
			name:     "Test Case 2: TypeFrontMatterORG",
			input:    TypeFrontMatterORG,
			expected: metadecoders.ORG,
		},
		{
			name:     "Test Case 3: TypeFrontMatterTOML",
			input:    TypeFrontMatterTOML,
			expected: metadecoders.TOML,
		},
		{
			name:     "Test Case 4: TypeFrontMatterYAML",
			input:    TypeFrontMatterYAML,
			expected: metadecoders.YAML,
		},
		{
			name:     "Test Case 5: Default Case",
			input:    0,
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

			result := FormatFromFrontMatterType(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
