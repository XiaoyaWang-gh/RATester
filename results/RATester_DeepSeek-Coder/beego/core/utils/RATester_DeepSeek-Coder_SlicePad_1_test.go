package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSlicePad_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		inputSlice []interface{}
		inputSize  int
		inputVal   interface{}
		expected   []interface{}
	}

	tests := []test{
		{[]interface{}{1, 2, 3}, 5, 0, []interface{}{1, 2, 3, 0, 0}},
		{[]interface{}{"a", "b", "c"}, 4, "x", []interface{}{"a", "b", "c", "x"}},
		{[]interface{}{}, 2, "y", []interface{}{"y", "y"}},
		{[]interface{}{1, 2, 3}, 2, 0, []interface{}{1, 2, 3}},
	}

	for _, test := range tests {
		result := SlicePad(test.inputSlice, test.inputSize, test.inputVal)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("SlicePad(%v, %v, %v) = %v; expected %v", test.inputSlice, test.inputSize, test.inputVal, result, test.expected)
		}
	}
}
