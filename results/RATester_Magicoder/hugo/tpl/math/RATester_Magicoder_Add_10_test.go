package math

import (
	"fmt"
	"testing"
)

func TestAdd_10(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name     string
		inputs   []any
		expected any
		err      error
	}{
		{
			name:     "add two integers",
			inputs:   []any{1, 2},
			expected: 3,
			err:      nil,
		},
		{
			name:     "add two floats",
			inputs:   []any{1.5, 2.5},
			expected: 4.0,
			err:      nil,
		},
		{
			name:     "add integer and float",
			inputs:   []any{1, 2.5},
			expected: 3.5,
			err:      nil,
		},
		{
			name:     "add two strings",
			inputs:   []any{"hello", "world"},
			expected: "helloworld",
			err:      nil,
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := ns.Add(test.inputs...)
			if err != test.err {
				t.Errorf("Expected error %v, but got %v", test.err, err)
			}
			if result != test.expected {
				t.Errorf("Expected result %v, but got %v", test.expected, result)
			}
		})
	}
}
