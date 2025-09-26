package strings

import (
	"errors"
	"fmt"
	"testing"
)

func TestToUpper_1(t *testing.T) {
	ns := Namespace{}

	tests := []struct {
		name     string
		input    any
		expected string
		err      error
	}{
		{
			name:     "Test with string",
			input:    "hello",
			expected: "HELLO",
			err:      nil,
		},
		{
			name:     "Test with int",
			input:    123,
			expected: "",
			err:      errors.New("unable to cast 123 of type int to string"),
		},
		{
			name:     "Test with nil",
			input:    nil,
			expected: "",
			err:      errors.New("unable to cast <nil> of type <nil> to string"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := ns.ToUpper(test.input)
			if err != nil && err.Error() != test.err.Error() {
				t.Errorf("Expected error %v, but got %v", test.err, err)
			}

			if result != test.expected {
				t.Errorf("Expected result %v, but got %v", test.expected, result)
			}
		})
	}
}
