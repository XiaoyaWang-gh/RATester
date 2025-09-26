package safe

import (
	"errors"
	"fmt"
	"html/template"
	"testing"
)

func TestCSS_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}

	type testCase struct {
		input    any
		expected template.CSS
		err      error
	}

	testCases := []testCase{
		{
			input:    "test",
			expected: "test",
			err:      nil,
		},
		{
			input:    123,
			expected: "123",
			err:      nil,
		},
		{
			input:    nil,
			expected: "",
			err:      nil,
		},
		{
			input:    true,
			expected: "",
			err:      errors.New("unable to cast true of type bool to string"),
		},
	}

	for _, tc := range testCases {
		result, err := ns.CSS(tc.input)
		if tc.err != nil {
			if err == nil {
				t.Errorf("Expected error, but got nil")
			} else if err.Error() != tc.err.Error() {
				t.Errorf("Expected error %v, but got %v", tc.err, err)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		}
	}
}
