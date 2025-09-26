package herrors

import (
	"errors"
	"fmt"
	"regexp"
	"testing"
)

func TestextractLineNo_1(t *testing.T) {
	re := regexp.MustCompile(`line (\d+)(?::(\d+))?`)
	extract := extractLineNo(re)

	tests := []struct {
		name     string
		input    error
		expected lineNumberExtractor
	}{
		{
			name:  "error with line number",
			input: errors.New("error at line 123"),
			expected: func(e error) (int, int) {
				return 123, 1
			},
		},
		{
			name:  "error with line and column number",
			input: errors.New("error at line 456:789"),
			expected: func(e error) (int, int) {
				return 456, 789
			},
		},
		{
			name:  "error without line number",
			input: errors.New("error at column 123"),
			expected: func(e error) (int, int) {
				return 0, 123
			},
		},
		{
			name:  "nil error",
			input: nil,
			expected: func(e error) (int, int) {
				panic("no error")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			lno, col := extract(test.input)
			expectedLno, expectedCol := test.expected(test.input)
			if lno != expectedLno || col != expectedCol {
				t.Errorf("expected (%d, %d), got (%d, %d)", expectedLno, expectedCol, lno, col)
			}
		})
	}
}
