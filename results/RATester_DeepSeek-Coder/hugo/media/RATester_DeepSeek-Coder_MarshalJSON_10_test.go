package media

import (
	"fmt"
	"testing"
)

func TestMarshalJSON_10(t *testing.T) {
	type testCase struct {
		name     string
		input    Type
		expected string
	}

	tests := []testCase{
		{
			name: "Test case 1",
			input: Type{
				Type:        "application/json",
				MainType:    "application",
				SubType:     "json",
				Delimiter:   ".",
				SuffixesCSV: "jpg,jpeg",
			},
			expected: `{"mainType":"application","subType":"json","delimiter":".","suffixes":["jpg","jpeg"]}`,
		},
		{
			name: "Test case 2",
			input: Type{
				Type:        "text/plain",
				MainType:    "text",
				SubType:     "plain",
				Delimiter:   "",
				SuffixesCSV: "",
			},
			expected: `{"mainType":"text","subType":"plain","delimiter":"","suffixes":[]}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := tc.input.MarshalJSON()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if string(result) != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, string(result))
			}
		})
	}
}
