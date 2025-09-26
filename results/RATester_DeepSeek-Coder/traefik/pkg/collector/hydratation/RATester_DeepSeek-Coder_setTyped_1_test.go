package hydratation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetTyped_1(t *testing.T) {
	type testStruct struct {
		field string
	}

	testCases := []struct {
		name     string
		field    reflect.Value
		i        interface{}
		expected reflect.Value
	}{
		{
			name:     "Test case 1",
			field:    reflect.ValueOf(testStruct{field: "oldValue"}.field),
			i:        "newValue",
			expected: reflect.ValueOf("newValue"),
		},
		{
			name:     "Test case 2",
			field:    reflect.ValueOf(testStruct{field: "oldValue"}.field),
			i:        123,
			expected: reflect.ValueOf("123"),
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
			if tc.field.String() != tc.expected.String() {
				t.Errorf("Expected %v, got %v", tc.expected.String(), tc.field.String())
			}
		})
	}
}
