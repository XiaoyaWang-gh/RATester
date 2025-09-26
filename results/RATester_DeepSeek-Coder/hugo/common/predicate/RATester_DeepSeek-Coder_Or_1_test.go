package predicate

import (
	"fmt"
	"reflect"
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
				return (i > 0) || (i < 0) || (i > 10)
			},
		},
		{
			name: "Test case 2",
			p: func(i int) bool {
				return i > 0
			},
			ps: []P[int]{
				func(i int) bool {
					return i > 0
				},
				func(i int) bool {
					return i > 10
				},
			},
			expected: func(i int) bool {
				return (i > 0) || (i > 10)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tt.p.Or(tt.ps...)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}
