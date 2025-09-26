package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBasicKind_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    reflect.Value
		expected kind
		err      error
	}{
		{
			name:     "bool",
			input:    reflect.ValueOf(true),
			expected: boolKind,
			err:      nil,
		},
		{
			name:     "int",
			input:    reflect.ValueOf(int(1)),
			expected: intKind,
			err:      nil,
		},
		{
			name:     "uint",
			input:    reflect.ValueOf(uint(1)),
			expected: uintKind,
			err:      nil,
		},
		{
			name:     "float",
			input:    reflect.ValueOf(float64(1.0)),
			expected: floatKind,
			err:      nil,
		},
		{
			name:     "complex",
			input:    reflect.ValueOf(complex(1.0, 1.0)),
			expected: complexKind,
			err:      nil,
		},
		{
			name:     "string",
			input:    reflect.ValueOf("test"),
			expected: stringKind,
			err:      nil,
		},
		{
			name:     "invalid",
			input:    reflect.ValueOf([]int{1, 2, 3}),
			expected: invalidKind,
			err:      errBadComparisonType,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := basicKind(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
			if err != tc.err {
				t.Errorf("Expected error %v, got %v", tc.err, err)
			}
		})
	}
}
