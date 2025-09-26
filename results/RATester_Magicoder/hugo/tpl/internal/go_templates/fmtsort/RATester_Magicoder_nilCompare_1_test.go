package fmtsort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestnilCompare_1(t *testing.T) {
	tests := []struct {
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
			bVal:     reflect.ValueOf(1),
			expected: -1,
			isEqual:  true,
		},
		{
			name:     "B is nil",
			aVal:     reflect.ValueOf(1),
			bVal:     reflect.ValueOf(nil),
			expected: 1,
			isEqual:  true,
		},
		{
			name:     "Neither nil",
			aVal:     reflect.ValueOf(1),
			bVal:     reflect.ValueOf(1),
			expected: 0,
			isEqual:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, isEqual := nilCompare(test.aVal, test.bVal)
			if result != test.expected || isEqual != test.isEqual {
				t.Errorf("Expected (%d, %t), got (%d, %t)", test.expected, test.isEqual, result, isEqual)
			}
		})
	}
}
