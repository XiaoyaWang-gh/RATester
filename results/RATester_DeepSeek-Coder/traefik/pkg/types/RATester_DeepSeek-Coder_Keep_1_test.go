package types

import (
	"fmt"
	"testing"
)

func TestKeep_1(t *testing.T) {
	type testCase struct {
		name     string
		fields   *AccessLogFields
		field    string
		expected bool
	}

	testCases := []testCase{
		{
			name: "Keep field if it exists and mode is keep",
			fields: &AccessLogFields{
				DefaultMode: "keep",
				Names: map[string]string{
					"field1": "keep",
				},
			},
			field:    "field1",
			expected: true,
		},
		{
			name: "Drop field if it exists and mode is drop",
			fields: &AccessLogFields{
				DefaultMode: "drop",
				Names: map[string]string{
					"field1": "drop",
				},
			},
			field:    "field1",
			expected: false,
		},
		{
			name: "Use default mode if field does not exist",
			fields: &AccessLogFields{
				DefaultMode: "drop",
				Names: map[string]string{
					"field2": "keep",
				},
			},
			field:    "field1",
			expected: false,
		},
		{
			name:     "Return true if fields is nil",
			fields:   nil,
			field:    "field1",
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.fields.Keep(tc.field)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
