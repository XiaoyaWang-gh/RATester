package types

import (
	"errors"
	"fmt"
	"testing"
)

func TestPrintableValue_4(t *testing.T) {
	t.Parallel()

	type testCase[T any] struct {
		name     string
		input    Result[T]
		expected any
	}

	testCases := []testCase[int]{
		{
			name: "success",
			input: Result[int]{
				Value: 1,
				Err:   nil,
			},
			expected: 1,
		},
		{
			name: "error",
			input: Result[int]{
				Value: 0,
				Err:   errors.New("test error"),
			},
			expected: "test error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			defer func() {
				if r := recover(); r != nil {
					if r.(error).Error() != tc.expected {
						t.Errorf("expected %v, got %v", tc.expected, r)
					}
				}
			}()

			result := tc.input.PrintableValue()
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}
