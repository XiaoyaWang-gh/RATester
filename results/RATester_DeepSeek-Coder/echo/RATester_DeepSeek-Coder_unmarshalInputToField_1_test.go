package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnmarshalInputToField_1(t *testing.T) {
	type TestStruct struct {
		Field string
	}

	testCases := []struct {
		name     string
		value    string
		field    reflect.Value
		expected bool
		err      error
	}{
		{
			name:     "Test case 1",
			value:    "test value",
			field:    reflect.ValueOf(&TestStruct{}).Elem().FieldByName("Field"),
			expected: true,
			err:      nil,
		},
		{
			name:     "Test case 2",
			value:    "test value",
			field:    reflect.ValueOf(&TestStruct{}).Elem().FieldByName("NonExistentField"),
			expected: false,
			err:      nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			actual, err := unmarshalInputToField(tc.field.Kind(), tc.value, tc.field)
			if actual != tc.expected || err != tc.err {
				t.Errorf("Expected (%v, %v), but got (%v, %v)", tc.expected, tc.err, actual, err)
			}
		})
	}
}
