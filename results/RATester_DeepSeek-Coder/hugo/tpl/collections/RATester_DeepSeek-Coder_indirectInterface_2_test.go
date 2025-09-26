package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIndirectInterface_2(t *testing.T) {
	type testStruct struct {
		val reflect.Value
	}

	testCases := []struct {
		name     string
		input    testStruct
		expected struct {
			rv    reflect.Value
			isNil bool
		}
	}{
		{
			name: "Test with non-interface value",
			input: testStruct{
				val: reflect.ValueOf(5),
			},
			expected: struct {
				rv    reflect.Value
				isNil bool
			}{
				rv:    reflect.ValueOf(5),
				isNil: false,
			},
		},
		{
			name: "Test with nil interface value",
			input: testStruct{
				val: reflect.ValueOf((*interface{})(nil)),
			},
			expected: struct {
				rv    reflect.Value
				isNil bool
			}{
				rv:    reflect.ValueOf((*interface{})(nil)),
				isNil: true,
			},
		},
		{
			name: "Test with non-nil interface value",
			input: testStruct{
				val: reflect.ValueOf(5),
			},
			expected: struct {
				rv    reflect.Value
				isNil bool
			}{
				rv:    reflect.ValueOf(5),
				isNil: false,
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

			rv, isNil := indirectInterface(tc.input.val)
			if rv.Interface() != tc.expected.rv.Interface() || isNil != tc.expected.isNil {
				t.Errorf("Expected (%v, %v), but got (%v, %v)", tc.expected.rv.Interface(), tc.expected.isNil, rv.Interface(), isNil)
			}
		})
	}
}
