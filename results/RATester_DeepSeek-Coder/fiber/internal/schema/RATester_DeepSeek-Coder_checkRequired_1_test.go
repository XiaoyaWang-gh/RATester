package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCheckRequired_1(t *testing.T) {
	type testStruct struct {
		Field1 string `required:"true"`
		Field2 string `required:"false"`
	}

	testCases := []struct {
		name     string
		input    map[string][]string
		expected MultiError
	}{
		{
			name: "all required fields are present",
			input: map[string][]string{
				"Field1": {"value1"},
				"Field2": {"value2"},
			},
			expected: MultiError{},
		},
		{
			name: "one required field is missing",
			input: map[string][]string{
				"Field2": {"value2"},
			},
			expected: MultiError{"Field1": EmptyFieldError{Key: "Field1"}},
		},
	}

	decoder := &Decoder{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := decoder.checkRequired(reflect.TypeOf(testStruct{}), tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
