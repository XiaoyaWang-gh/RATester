package strings

import (
	"errors"
	"fmt"
	"testing"
)

func TestTrim_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name     string
		s        any
		cutset   any
		expected string
		err      error
	}{
		{
			name:     "trim string",
			s:        "  hello world !  ",
			cutset:   " !",
			expected: "hello world",
			err:      nil,
		},
		{
			name:     "trim string with error",
			s:        "  hello world !  ",
			cutset:   123,
			expected: "",
			err:      errors.New("unable to cast 123 of type int to string"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := ns.Trim(test.s, test.cutset)
			if err != nil && err.Error() != test.err.Error() {
				t.Errorf("Expected error %v, got %v", test.err, err)
			}

			if result != test.expected {
				t.Errorf("Expected result %s, got %s", test.expected, result)
			}
		})
	}
}
