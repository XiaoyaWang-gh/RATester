package strings

import (
	"errors"
	"fmt"
	"testing"
)

func TestContainsAny_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name     string
		s        any
		chars    any
		expected bool
		err      error
	}{
		{
			name:     "contains any",
			s:        "hello",
			chars:    "eo",
			expected: true,
			err:      nil,
		},
		{
			name:     "does not contain any",
			s:        "hello",
			chars:    "xyz",
			expected: false,
			err:      nil,
		},
		{
			name:     "error case",
			s:        "hello",
			chars:    123,
			expected: false,
			err:      errors.New("unsupported type"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := ns.ContainsAny(test.s, test.chars)
			if err != nil && err.Error() != test.err.Error() {
				t.Errorf("Expected error %v, but got %v", test.err, err)
			}
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
