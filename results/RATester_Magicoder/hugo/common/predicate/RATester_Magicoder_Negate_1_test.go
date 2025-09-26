package predicate

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNegate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type testCase[T any] struct {
		input    P[T]
		expected P[T]
	}

	testCases := []testCase[int]{
		{
			input: func(v int) bool {
				return v > 0
			},
			expected: func(v int) bool {
				return v <= 0
			},
		},
		{
			input: func(v int) bool {
				return v < 0
			},
			expected: func(v int) bool {
				return v >= 0
			},
		},
	}

	for _, tc := range testCases {
		result := tc.input.Negate()
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v, but got %v", tc.expected, result)
		}
	}
}
