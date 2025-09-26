package gin

import (
	"fmt"
	"testing"
)

func TestIsType_1(t *testing.T) {
	testCases := []struct {
		name     string
		err      *Error
		flags    ErrorType
		expected bool
	}{
		{
			name: "Test case 1",
			err: &Error{
				Type: ErrorType(1),
			},
			flags:    ErrorType(1),
			expected: true,
		},
		{
			name: "Test case 2",
			err: &Error{
				Type: ErrorType(1),
			},
			flags:    ErrorType(2),
			expected: false,
		},
		{
			name: "Test case 3",
			err: &Error{
				Type: ErrorType(3),
			},
			flags:    ErrorType(1 | 2),
			expected: true,
		},
		{
			name: "Test case 4",
			err: &Error{
				Type: ErrorType(0),
			},
			flags:    ErrorType(1 | 2),
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.err.IsType(tc.flags)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
