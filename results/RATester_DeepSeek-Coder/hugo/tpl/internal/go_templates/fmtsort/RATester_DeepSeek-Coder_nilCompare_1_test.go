package fmtsort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNilCompare_1(t *testing.T) {
	testCases := []struct {
		name     string
		aVal     reflect.Value
		bVal     reflect.Value
		expected int
		isEqual  bool
	}{
		{
			name:     "Both nil",
			aVal:     reflect.ValueOf(nil),
			bVal:     reflect.ValueOf(nil),
			expected: 0,
			isEqual:  true,
		},
		{
			name:     "A is nil",
			aVal:     reflect.ValueOf(nil),
			bVal:     reflect.ValueOf("not nil"),
			expected: -1,
			isEqual:  true,
		},
		{
			name:     "B is nil",
			aVal:     reflect.ValueOf("not nil"),
			bVal:     reflect.ValueOf(nil),
			expected: 1,
			isEqual:  true,
		},
		{
			name:     "Neither is nil",
			aVal:     reflect.ValueOf("not nil"),
			bVal:     reflect.ValueOf("also not nil"),
			expected: 0,
			isEqual:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, isEqual := nilCompare(tc.aVal, tc.bVal)
			if result != tc.expected || isEqual != tc.isEqual {
				t.Errorf("Expected (%d, %t), but got (%d, %t)", tc.expected, tc.isEqual, result, isEqual)
			}
		})
	}
}
