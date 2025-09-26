package strings

import (
	"errors"
	"fmt"
	"testing"
)

func TestTrimLeft_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name     string
		cutset   any
		s        any
		expected string
		err      error
	}{
		{
			name:     "TrimLeft with string",
			cutset:   "abc",
			s:        "abcdefg",
			expected: "defg",
			err:      nil,
		},
		{
			name:     "TrimLeft with non-string cutset",
			cutset:   123,
			s:        "abcdefg",
			expected: "",
			err:      errors.New("cast: unable to cast \"123\" of type int to string"),
		},
		{
			name:     "TrimLeft with non-string s",
			cutset:   "abc",
			s:        123,
			expected: "",
			err:      errors.New("cast: unable to cast \"123\" of type int to string"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := ns.TrimLeft(test.cutset, test.s)
			if err != nil && err.Error() != test.err.Error() {
				t.Errorf("Expected error %v, but got %v", test.err, err)
			}
			if result != test.expected {
				t.Errorf("Expected result %v, but got %v", test.expected, result)
			}
		})
	}
}
