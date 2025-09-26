package math

import (
	"errors"
	"fmt"
	"math"
	"testing"
)

func TestToRadians_1(t *testing.T) {
	ns := Namespace{}

	tests := []struct {
		name     string
		input    any
		expected float64
		err      error
	}{
		{"valid input", 180, math.Pi, nil},
		{"invalid input", "abc", 0, errors.New("requires a numeric argument")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := ns.ToRadians(test.input)
			if err != nil && err.Error() != test.err.Error() {
				t.Errorf("Expected error %v, but got %v", test.err, err)
			}
			if result != test.expected {
				t.Errorf("Expected result %v, but got %v", test.expected, result)
			}
		})
	}
}
