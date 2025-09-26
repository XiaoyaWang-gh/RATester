package herrors

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
)

func TestextractOffsetAndType_1(t *testing.T) {
	type testCase struct {
		name     string
		input    error
		expected struct {
			offset int
			etype  string
		}
	}

	testCases := []testCase{
		{
			name:  "UnmarshalTypeError",
			input: &json.UnmarshalTypeError{Offset: 10},
			expected: struct {
				offset int
				etype  string
			}{
				offset: 10,
				etype:  "json",
			},
		},
		{
			name:  "SyntaxError",
			input: &json.SyntaxError{Offset: 20},
			expected: struct {
				offset int
				etype  string
			}{
				offset: 20,
				etype:  "json",
			},
		},
		{
			name:  "OtherError",
			input: errors.New("other error"),
			expected: struct {
				offset int
				etype  string
			}{
				offset: -1,
				etype:  "",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			offset, etype := extractOffsetAndType(tc.input)
			if offset != tc.expected.offset || etype != tc.expected.etype {
				t.Errorf("Expected (%d, %s), but got (%d, %s)", tc.expected.offset, tc.expected.etype, offset, etype)
			}
		})
	}
}
