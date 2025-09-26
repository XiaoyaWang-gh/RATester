package types

import (
	"fmt"
	"testing"
)

func TestKeepHeader_1(t *testing.T) {
	type testCase struct {
		name     string
		fields   *AccessLogFields
		header   string
		expected string
	}

	testCases := []testCase{
		{
			name: "Header exists in the map",
			fields: &AccessLogFields{
				Headers: &FieldHeaders{
					Names: map[string]string{"header1": "keep"},
				},
			},
			header:   "header1",
			expected: "keep",
		},
		{
			name: "Header does not exist in the map",
			fields: &AccessLogFields{
				Headers: &FieldHeaders{
					Names: map[string]string{"header2": "drop"},
				},
			},
			header:   "header1",
			expected: "drop",
		},
		{
			name: "Header exists in the map with default mode",
			fields: &AccessLogFields{
				DefaultMode: "keep",
				Headers: &FieldHeaders{
					Names: map[string]string{"header1": "drop"},
				},
			},
			header:   "header1",
			expected: "drop",
		},
		{
			name: "Header does not exist in the map with default mode",
			fields: &AccessLogFields{
				DefaultMode: "drop",
				Headers: &FieldHeaders{
					Names: map[string]string{"header2": "keep"},
				},
			},
			header:   "header1",
			expected: "keep",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.fields.KeepHeader(tc.header)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
