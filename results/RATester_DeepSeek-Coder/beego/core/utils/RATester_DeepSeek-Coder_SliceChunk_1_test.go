package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSliceChunk_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		inputSlice []interface{}
		inputSize  int
		expected   [][]interface{}
	}

	tests := []test{
		{
			inputSlice: []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			inputSize:  2,
			expected:   [][]interface{}{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}},
		},
		{
			inputSlice: []interface{}{"a", "b", "c", "d", "e"},
			inputSize:  2,
			expected:   [][]interface{}{{"a", "b"}, {"c", "d"}, {"e"}},
		},
		{
			inputSlice: []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			inputSize:  10,
			expected:   [][]interface{}{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		},
		{
			inputSlice: []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			inputSize:  11,
			expected:   [][]interface{}{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		},
	}

	for _, test := range tests {
		result := SliceChunk(test.inputSlice, test.inputSize)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
	}
}
