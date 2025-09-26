package math

import (
	"fmt"
	"testing"
)

func TestMax_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name     string
		inputs   []any
		expected float64
		err      error
	}{
		{
			name:     "single input",
			inputs:   []any{1.0},
			expected: 1.0,
			err:      nil,
		},
		{
			name:     "multiple inputs",
			inputs:   []any{1.0, 2.0, 3.0},
			expected: 3.0,
			err:      nil,
		},
		{
			name:     "negative numbers",
			inputs:   []any{-1.0, -2.0, -3.0},
			expected: -1.0,
			err:      nil,
		},
		{
			name:     "mixed numbers",
			inputs:   []any{-1.0, 2.0, -3.0},
			expected: 2.0,
			err:      nil,
		},
		{
			name:     "zero",
			inputs:   []any{0.0, 0.0, 0.0},
			expected: 0.0,
			err:      nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			maximum, err := ns.Max(test.inputs...)
			if err != test.err {
				t.Errorf("Expected error %v, but got %v", test.err, err)
			}
			if maximum != test.expected {
				t.Errorf("Expected maximum %v, but got %v", test.expected, maximum)
			}
		})
	}
}
