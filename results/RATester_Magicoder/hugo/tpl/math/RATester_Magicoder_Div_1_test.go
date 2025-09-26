package math

import (
	"errors"
	"fmt"
	"testing"
)

func TestDiv_1(t *testing.T) {
	ns := &Namespace{}

	testCases := []struct {
		name     string
		inputs   []any
		expected any
		err      error
	}{
		{
			name:     "Divide two numbers",
			inputs:   []any{10, 2},
			expected: 5.0,
			err:      nil,
		},
		{
			name:     "Divide by zero",
			inputs:   []any{10, 0},
			expected: nil,
			err:      errors.New("division by zero"),
		},
		{
			name:     "Divide by string",
			inputs:   []any{10, "2"},
			expected: nil,
			err:      errors.New("unsupported operand type"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := ns.Div(tc.inputs...)
			if err != nil && err.Error() != tc.err.Error() {
				t.Errorf("Expected error %v, but got %v", tc.err, err)
			}
			if result != tc.expected {
				t.Errorf("Expected result %v, but got %v", tc.expected, result)
			}
		})
	}
}
