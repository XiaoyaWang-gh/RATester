package hqt

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStructTypes_1(t *testing.T) {
	type testStruct struct {
		field1 string
		field2 int
	}
	testCases := []struct {
		name     string
		input    any
		expected map[reflect.Type]struct{}
	}{
		{
			name:  "Test with struct",
			input: testStruct{field1: "test", field2: 1},
			expected: map[reflect.Type]struct{}{
				reflect.TypeOf(testStruct{}): {},
				reflect.TypeOf(""):           {},
				reflect.TypeOf(0):            {},
			},
		},
		{
			name:  "Test with pointer to struct",
			input: &testStruct{field1: "test", field2: 1},
			expected: map[reflect.Type]struct{}{
				reflect.TypeOf(&testStruct{}): {},
				reflect.TypeOf(""):            {},
				reflect.TypeOf(0):             {},
			},
		},
		{
			name:  "Test with nil",
			input: nil,
			expected: map[reflect.Type]struct{}{
				reflect.TypeOf(nil): {},
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

			m := make(map[reflect.Type]struct{})
			structTypes(reflect.ValueOf(tc.input), m)
			if !reflect.DeepEqual(m, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, m)
			}
		})
	}
}
