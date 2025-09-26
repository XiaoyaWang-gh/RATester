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
			name:  "empty map",
			input: map[string]int{},
			expected: SortedMap{
				{Key: reflect.ValueOf(""), Value: reflect.ValueOf(0)},
			},
		},
		{
			name:  "single element",
			input: map[string]int{"a": 1},
			expected: SortedMap{
				{Key: reflect.ValueOf("a"), Value: reflect.ValueOf(1)},
			},
		},
		{
			name: "multiple elements",
			input: map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			expected: SortedMap{
				{Key: reflect.ValueOf("a"), Value: reflect.ValueOf(1)},
				{Key: reflect.ValueOf("b"), Value: reflect.ValueOf(2)},
				{Key: reflect.ValueOf("c"), Value: reflect.ValueOf(3)},
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
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}
