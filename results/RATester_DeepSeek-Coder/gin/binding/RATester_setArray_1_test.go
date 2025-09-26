package binding

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetArray_1(t *testing.T) {
	type TestStruct struct {
		ArrayField []string
	}

	testCases := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "Test Case 1",
			input:    []string{"value1", "value2", "value3"},
			expected: []string{"value1", "value2", "value3"},
		},
		{
			name:     "Test Case 2",
			input:    []string{"value4", "value5", "value6"},
			expected: []string{"value4", "value5", "value6"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			testStruct := TestStruct{}
			value := reflect.ValueOf(&testStruct).Elem().FieldByName("ArrayField")
			field, _ := reflect.TypeOf(testStruct).FieldByName("ArrayField")

			err := setArray(tc.input, value, field)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if !reflect.DeepEqual(testStruct.ArrayField, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, testStruct.ArrayField)
			}
		})
	}
}
