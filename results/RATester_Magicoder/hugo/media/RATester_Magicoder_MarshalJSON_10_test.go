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

	testCases := []testCase{
		{
			name: "Test Case 1",
			input: Type{
				Type:        "application/json",
				MainType:    "application",
				SubType:     "json",
				Delimiter:   ".",
				FirstSuffix: SuffixInfo{Suffix: "json", FullSuffix: ".json"},
				SuffixesCSV: "json,xml",
			},
			expected: `{"mainType":"application","subType":"json","delimiter":".","type":"application/json","string":"application/json","suffixes":["json","xml"]}`,
		},
		{
			name: "Test Case 2",
			input: Type{
				Type:        "text/plain",
				MainType:    "text",
				SubType:     "plain",
				Delimiter:   "",
				FirstSuffix: SuffixInfo{Suffix: "plain", FullSuffix: "plain"},
				SuffixesCSV: "txt,log",
			},
			expected: `{"mainType":"text","subType":"plain","delimiter":"","type":"text/plain","string":"text/plain","suffixes":["txt","log"]}`,
		},
	}

	for _, tc := range testCases {
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
				t.Errorf("Expected: %s, got: %s", tc.expected, string(result))
			}
		})
	}
}
