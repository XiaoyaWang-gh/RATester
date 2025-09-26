package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDrain_3(t *testing.T) {
	t.Parallel()

	type testCase[T any] struct {
		name   string
		stack  *Stack[T]
		expect []T
	}

	tests := []testCase[int]{
		{
			name:   "empty stack",
			stack:  &Stack[int]{},
			expect: []int{},
		},
		{
			name:   "single item",
			stack:  &Stack[int]{items: []int{1}},
			expect: []int{1},
		},
		{
			name:   "multiple items",
			stack:  &Stack[int]{items: []int{1, 2, 3}},
			expect: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tt.stack.Drain()
			if !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("expected %v, got %v", tt.expect, got)
			}
		})
	}
}
