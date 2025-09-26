package hydratation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetTyped_1(t *testing.T) {
	type testStruct struct {
		Field1 int
		Field2 string
	}

	testCases := []struct {
		name     string
		field    reflect.Value
		i        interface{}
		expected reflect.Value
	}{
		{
			name:     "Set int field",
			field:    reflect.ValueOf(&testStruct{}).Elem().FieldByName("Field1"),
			i:        10,
			expected: reflect.ValueOf(10),
		},
		{
			name:     "Set string field",
			field:    reflect.ValueOf(&testStruct{}).Elem().FieldByName("Field2"),
			i:        "test",
			expected: reflect.ValueOf("test"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			setTyped(tc.field, tc.i)
			if tc.field.Interface() != tc.expected.Interface() {
				t.Errorf("Expected %v, but got %v", tc.expected.Interface(), tc.field.Interface())
			}
		})
	}
}
