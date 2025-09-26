package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAppendToInterfaceSliceFromValues_1(t *testing.T) {
	type testCase struct {
		name     string
		slice1   reflect.Value
		slice2   reflect.Value
		expected []any
	}

	testCases := []testCase{
		{
			name:     "both slices are valid",
			slice1:   reflect.ValueOf([]int{1, 2, 3}),
			slice2:   reflect.ValueOf([]string{"a", "b", "c"}),
			expected: []any{1, 2, 3, "a", "b", "c"},
		},
		{
			name:     "slice1 is valid, slice2 is invalid",
			slice1:   reflect.ValueOf([]int{1, 2, 3}),
			slice2:   reflect.Value{},
			expected: []any{1, 2, 3, nil},
		},
		{
			name:     "slice1 is invalid, slice2 is valid",
			slice1:   reflect.Value{},
			slice2:   reflect.ValueOf([]string{"a", "b", "c"}),
			expected: []any{nil, "a", "b", "c"},
		},
		{
			name:     "both slices are invalid",
			slice1:   reflect.Value{},
			slice2:   reflect.Value{},
			expected: []any{nil, nil},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := appendToInterfaceSliceFromValues(tc.slice1, tc.slice2)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
