package predicate

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNegate_1(t *testing.T) {
	type testCase[T any] struct {
		name     string
		p        P[T]
		expected P[T]
	}

	testCases := []testCase[int]{
		{
			name: "Test case 1",
			p: func(i int) bool {
				return i > 0
			},
			expected: func(i int) bool {
				return !(i > 0)
			},
		},
		{
			name: "Test case 2",
			p: func(i int) bool {
				return i < 0
			},
			expected: func(i int) bool {
				return !(i < 0)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.p.Negate()
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
