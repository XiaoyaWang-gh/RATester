package herrors

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
)

func TestExtractOffsetAndType_1(t *testing.T) {
	type testCase struct {
		name     string
		input    error
		expected struct {
			offset int
			typ    string
		}
	}

	testCases := []testCase{
		{
			name:  "UnmarshalTypeError",
			input: &json.UnmarshalTypeError{Offset: 10},
			expected: struct {
				offset int
				typ    string
			}{offset: 10, typ: "json"},
		},
		{
			name:  "SyntaxError",
			input: &json.SyntaxError{Offset: 20},
			expected: struct {
				offset int
				typ    string
			}{offset: 20, typ: "json"},
		},
		{
			name:  "Default",
			input: errors.New("unknown error"),
			expected: struct {
				offset int
				typ    string
			}{offset: -1, typ: ""},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			offset, typ := extractOffsetAndType(tc.input)
			if offset != tc.expected.offset || typ != tc.expected.typ {
				t.Errorf("Expected (%d, %s) but got (%d, %s)", tc.expected.offset, tc.expected.typ, offset, typ)
			}
		})
	}
}
