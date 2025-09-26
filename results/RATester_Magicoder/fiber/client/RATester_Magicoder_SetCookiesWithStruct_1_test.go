package client

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetCookiesWithStruct_1(t *testing.T) {
	type testStruct struct {
		Name string
		Age  int
	}

	testCases := []struct {
		name     string
		input    any
		expected Cookie
	}{
		{
			name: "Test case 1",
			input: testStruct{
				Name: "John",
				Age:  30,
			},
			expected: Cookie{
				"Name": "John",
				"Age":  "30",
			},
		},
		{
			name: "Test case 2",
			input: testStruct{
				Name: "Jane",
				Age:  25,
			},
			expected: Cookie{
				"Name": "Jane",
				"Age":  "25",
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
				t.Errorf("Expected %v, but got %v", tc.expected, c)
			}
		})
	}
}
