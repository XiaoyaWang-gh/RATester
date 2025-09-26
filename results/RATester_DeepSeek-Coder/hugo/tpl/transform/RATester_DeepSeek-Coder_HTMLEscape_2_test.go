package transform

import (
	"errors"
	"fmt"
	"testing"
)

func TestHTMLEscape_2(t *testing.T) {
	type testCase struct {
		name     string
		input    any
		expected string
		err      error
	}

	tests := []testCase{
		{
			name:     "String input",
			input:    "<div>Hello, World</div>",
			expected: "&lt;div&gt;Hello, World&lt;/div&gt;",
			err:      nil,
		},
		{
			name:     "Non-string input",
			input:    123,
			expected: "",
			err:      errors.New("cast: unable to cast 123 of type int to string"),
		},
	}

	ns := &Namespace{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := ns.HTMLEscape(tc.input)
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
					t.Errorf("Expected %q, but got %q", tc.expected, result)
				}
			}
		})
	}
}
