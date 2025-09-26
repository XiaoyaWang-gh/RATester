package media

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSuffixes_1(t *testing.T) {
	type testCase struct {
		name     string
		input    Type
		expected []string
	}

	testCases := []testCase{
		{
			name: "empty suffixes",
			input: Type{
				SuffixesCSV: "",
			},
			expected: nil,
		},
		{
			name: "single suffix",
			input: Type{
				SuffixesCSV: "jpg",
			},
			expected: []string{"jpg"},
		},
		{
			name: "multiple suffixes",
			input: Type{
				SuffixesCSV: "jpg,jpeg",
			},
			expected: []string{"jpg", "jpeg"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.input.Suffixes()
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, got: %v", tc.expected, result)
			}
		})
	}
}
