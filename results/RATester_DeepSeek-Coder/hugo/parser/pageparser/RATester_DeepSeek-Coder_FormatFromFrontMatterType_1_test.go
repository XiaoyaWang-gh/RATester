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
			name:     "Test JSON",
			input:    TypeFrontMatterJSON,
			expected: metadecoders.JSON,
		},
		{
			name:     "Test ORG",
			input:    TypeFrontMatterORG,
			expected: metadecoders.ORG,
		},
		{
			name:     "Test TOML",
			input:    TypeFrontMatterTOML,
			expected: metadecoders.TOML,
		},
		{
			name:     "Test YAML",
			input:    TypeFrontMatterYAML,
			expected: metadecoders.YAML,
		},
		{
			name:     "Test Unknown",
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
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
