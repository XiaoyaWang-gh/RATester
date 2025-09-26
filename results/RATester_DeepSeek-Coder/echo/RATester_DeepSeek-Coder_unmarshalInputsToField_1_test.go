package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnmarshalInputsToField_1(t *testing.T) {
	type TestStruct struct {
		Field string
	}

	testCases := []struct {
		name        string
		valueKind   reflect.Kind
		values      []string
		field       reflect.Value
		expected    bool
		expectedErr error
	}{
		{
			name:      "Test case 1",
			valueKind: reflect.Ptr,
			values:    []string{"test"},
			field:     reflect.ValueOf(&TestStruct{}),
			expected:  true,
		},
		{
			name:      "Test case 2",
			valueKind: reflect.Struct,
			values:    []string{"test"},
			field:     reflect.ValueOf(TestStruct{}),
			expected:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			actual, err := unmarshalInputsToField(tc.valueKind, tc.values, tc.field)
			if (err != nil) != (tc.expectedErr != nil) || actual != tc.expected {
				t.Errorf("Expected (%v, %v), but got (%v, %v)", tc.expected, tc.expectedErr, actual, err)
			}
		})
	}
}
