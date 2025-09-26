package context

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFormTagName_1(t *testing.T) {
	type TestStruct struct {
		Field1 string `form:"field1"`
		Field2 string `form:"-"`
		Field3 string `form:""`
		Field4 string `form:"field4,omitempty"`
	}

	testCases := []struct {
		name     string
		field    reflect.StructField
		expected string
		ok       bool
	}{
		{
			name:     "Test Case 1",
			field:    reflect.TypeOf(TestStruct{}).Field(0),
			expected: "field1",
			ok:       true,
		},
		{
			name:     "Test Case 2",
			field:    reflect.TypeOf(TestStruct{}).Field(1),
			expected: "",
			ok:       false,
		},
		{
			name:     "Test Case 3",
			field:    reflect.TypeOf(TestStruct{}).Field(2),
			expected: "Field3",
			ok:       true,
		},
		{
			name:     "Test Case 4",
			field:    reflect.TypeOf(TestStruct{}).Field(3),
			expected: "field4",
			ok:       true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			actual, ok := formTagName(tc.field)
			if actual != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, actual)
			}
			if ok != tc.ok {
				t.Errorf("Expected %v, but got %v", tc.ok, ok)
			}
		})
	}
}
