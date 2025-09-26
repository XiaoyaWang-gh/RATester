package binder

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseToStruct_1(t *testing.T) {
	type TestStruct struct {
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
	}

	testCases := []struct {
		name     string
		aliasTag string
		out      any
		data     map[string][]string
		expected any
	}{
		{
			name:     "Test Case 1",
			aliasTag: "json",
			out:      &TestStruct{},
			data:     map[string][]string{"field1": {"value1"}, "field2": {"123"}},
			expected: &TestStruct{Field1: "value1", Field2: 123},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := parseToStruct(tc.aliasTag, tc.out, tc.data)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if !reflect.DeepEqual(tc.out, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, tc.out)
			}
		})
	}
}
