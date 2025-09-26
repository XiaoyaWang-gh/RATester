package predicate

import (
	"fmt"
	"testing"
)

func TestOr_1(t *testing.T) {
	type testCase[T any] struct {
		name     string
		p        P[T]
		ps       []P[T]
		expected P[T]
	}

	tests := []testCase[int]{
		{
			name: "Test case 1",
			p: func(i int) bool {
				return i > 0
			},
			ps: []P[int]{
				func(i int) bool {
					return i < 0
				},
				func(i int) bool {
					return i > 10
				},
			},
			expected: func(i int) bool {
				return i > 0 || i < 0 || i > 10
			},
		},
		{
			name: "Test case 2",
			p: func(i int) bool {
				return i == 0
			},
			ps: []P[int]{
				func(i int) bool {
					return i == 1
				},
				func(i int) bool {
					return i == 2
				},
			},
			expected: func(i int) bool {
				return i == 0 || i == 1 || i == 2
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.p.Or(tc.ps...)
			for i := -10; i <= 10; i++ {
				if result(i) != tc.expected(i) {
					t.Errorf("Expected %v, but got %v for input %v", tc.expected(i), result(i), i)
				}
			}
		})
	}
}
