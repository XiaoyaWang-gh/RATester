package strings

import (
	"errors"
	"fmt"
	"testing"
)

func TestTrimPrefix_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name     string
		prefix   any
		s        any
		expected string
		err      error
	}{
		{
			name:     "TrimPrefix",
			prefix:   "prefix",
			s:        "prefix_string",
			expected: "string",
			err:      nil,
		},
		{
			name:     "TrimPrefix with non-string prefix",
			prefix:   123,
			s:        "prefix_string",
			expected: "",
			err:      errors.New("cast: unable to cast 123 of type int to string"),
		},
		{
			name:     "TrimPrefix with non-string s",
			prefix:   "prefix",
			s:        123,
			expected: "",
			err:      errors.New("cast: unable to cast 123 of type int to string"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := ns.TrimPrefix(test.prefix, test.s)
			if err != nil && err.Error() != test.err.Error() {
				t.Errorf("Expected error %v, but got %v", test.err, err)
			}

			if result != test.expected {
				t.Errorf("Expected result %v, but got %v", test.expected, result)
			}
		})
	}
}
