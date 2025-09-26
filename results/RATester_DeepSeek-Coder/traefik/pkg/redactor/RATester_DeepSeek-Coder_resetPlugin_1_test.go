package redactor

import (
	"fmt"
	"reflect"
	"testing"
)

func TestResetPlugin_1(t *testing.T) {
	type testStruct struct {
		field1 string
		field2 int
	}

	testCases := []struct {
		name     string
		input    reflect.Value
		expected reflect.Value
	}{
		{
			name:     "Test Case 1",
			input:    reflect.ValueOf(testStruct{"test", 1}),
			expected: reflect.ValueOf(testStruct{"", 0}),
		},
		{
			name:     "Test Case 2",
			input:    reflect.ValueOf(testStruct{"", 0}),
			expected: reflect.ValueOf(testStruct{"", 0}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			resetPlugin(tc.input)
			if !reflect.DeepEqual(tc.input.Interface(), tc.expected.Interface()) {
				t.Errorf("Expected %v, got %v", tc.expected.Interface(), tc.input.Interface())
			}
		})
	}
}
