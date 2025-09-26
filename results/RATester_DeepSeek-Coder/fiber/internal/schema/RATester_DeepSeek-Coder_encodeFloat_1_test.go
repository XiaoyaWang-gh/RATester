package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEncodeFloat_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    reflect.Value
		expected string
	}{
		{
			name:     "Test with float32",
			input:    reflect.ValueOf(float32(123.456)),
			expected: "123.456000",
		},
		{
			name:     "Test with float64",
			input:    reflect.ValueOf(float64(123.456)),
			expected: "123.456000",
		},
		{
			name:     "Test with complex64",
			input:    reflect.ValueOf(complex64(123.456 + 789.101i)),
			expected: "(123.456+789.101i)",
		},
		{
			name:     "Test with complex128",
			input:    reflect.ValueOf(complex128(123.456 + 789.101i)),
			expected: "(123.456+789.101i)",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := encodeFloat(tc.input, 6)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
