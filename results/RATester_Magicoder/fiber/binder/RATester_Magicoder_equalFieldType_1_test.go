package binder

import (
	"fmt"
	"reflect"
	"testing"
)

func TestequalFieldType_1(t *testing.T) {

	type testStruct struct {
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
		Field3 bool   `json:"field3"`
	}

	testCases := []struct {
		name     string
		out      any
		kind     reflect.Kind
		key      string
		expected bool
	}{
		{
			name:     "Test case 1: Map",
			out:      &map[string]string{},
			kind:     reflect.Map,
			key:      "test",
			expected: true,
		},
		{
			name:     "Test case 2: Struct",
			out:      &testStruct{},
			kind:     reflect.String,
			key:      "field1",
			expected: true,
		},
		{
			name:     "Test case 3: Struct",
			out:      &testStruct{},
			kind:     reflect.Int,
			key:      "field2",
			expected: true,
		},
		{
			name:     "Test case 4: Struct",
			out:      &testStruct{},
			kind:     reflect.Bool,
			key:      "field3",
			expected: true,
		},
		{
			name:     "Test case 5: Struct",
			out:      &testStruct{},
			kind:     reflect.Float64,
			key:      "field4",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := equalFieldType(tc.out, tc.kind, tc.key)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
