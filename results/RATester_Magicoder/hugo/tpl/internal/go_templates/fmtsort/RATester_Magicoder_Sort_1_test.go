package fmtsort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSort_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    map[string]int
		expected SortedMap
	}{
		{
			name:     "empty map",
			input:    map[string]int{},
			expected: SortedMap{},
		},
		{
			name:  "single element",
			input: map[string]int{"a": 1},
			expected: SortedMap{
				{reflect.ValueOf("a"), reflect.ValueOf(1)},
			},
		},
		{
			name:  "multiple elements",
			input: map[string]int{"a": 1, "b": 2, "c": 3},
			expected: SortedMap{
				{reflect.ValueOf("a"), reflect.ValueOf(1)},
				{reflect.ValueOf("b"), reflect.ValueOf(2)},
				{reflect.ValueOf("c"), reflect.ValueOf(3)},
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

			mapValue := reflect.ValueOf(tc.input)
			result := Sort(mapValue)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got: %v", tc.expected, result)
			}
		})
	}
}
