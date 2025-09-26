package math

import (
	"errors"
	"fmt"
	"testing"
)

func TestCeil_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name     string
		input    any
		expected float64
		err      error
	}{
		{"positive integer", 1, 1, nil},
		{"positive float", 1.5, 2, nil},
		{"negative integer", -1, -1, nil},
		{"negative float", -1.5, -1, nil},
		{"zero", 0, 0, nil},
		{"non-float value", "string", 0, errors.New("Ceil operator can't be used with non-float value")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := ns.Ceil(test.input)
			if err != nil && err.Error() != test.err.Error() {
				t.Errorf("Expected error %v, but got %v", test.err, err)
			}
			if result != test.expected {
				t.Errorf("Expected result %v, but got %v", test.expected, result)
			}
		})
	}
}
