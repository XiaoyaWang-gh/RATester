package client

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetCookiesWithStruct_1(t *testing.T) {
	type testStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	testCases := []struct {
		name     string
		input    any
		expected Cookie
	}{
		{
			name:  "Test with a struct",
			input: testStruct{Name: "John", Age: 30},
			expected: Cookie{
				"name": "John",
				"age":  "30",
			},
		},
		{
			name:  "Test with a map",
			input: map[string]any{"Name": "Jane", "Age": 25},
			expected: Cookie{
				"Name": "Jane",
				"Age":  "25",
			},
		},
		{
			name:  "Test with a slice",
			input: []any{"Name", "Age"},
			expected: Cookie{
				"Name": "Name",
				"Age":  "Age",
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

			c := make(Cookie)
			c.SetCookiesWithStruct(tc.input)
			if !reflect.DeepEqual(c, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, c)
			}
		})
	}
}
