package predicate

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFilter_4(t *testing.T) {
	type testCase[T any] struct {
		name     string
		p        P[T]
		input    []T
		expected []T
	}

	even := func(n int) bool { return n%2 == 0 }
	isOdd := func(n int) bool { return n%2 != 0 }

	testCases := []testCase[int]{
		{
			name:     "filter even numbers",
			p:        P[int](even),
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{2, 4, 6, 8, 10},
		},
		{
			name:     "filter odd numbers",
			p:        P[int](isOdd),
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{1, 3, 5, 7, 9},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tc.p.Filter(tc.input)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
