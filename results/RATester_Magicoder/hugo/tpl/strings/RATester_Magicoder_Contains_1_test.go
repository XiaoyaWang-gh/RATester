package strings

import (
	"errors"
	"fmt"
	"testing"
)

func TestContains_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}

	tests := []struct {
		s        any
		substr   any
		expected bool
		err      error
	}{
		{"hello", "lo", true, nil},
		{"hello", "world", false, nil},
		{123, "23", false, errors.New("unable to cast 123 of type int to string")},
		{"hello", 123, false, errors.New("unable to cast 123 of type int to string")},
	}

	for _, test := range tests {
		result, err := ns.Contains(test.s, test.substr)
		if err != nil && err.Error() != test.err.Error() {
			t.Errorf("Expected error %v, but got %v", test.err, err)
		}
		if result != test.expected {
			t.Errorf("Expected %v, but got %v", test.expected, result)
		}
	}
}
