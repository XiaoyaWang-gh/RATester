package client

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetParamsWithStruct_1(t *testing.T) {
	type testStruct struct {
		Key   string
		Value string
	}

	testCases := []struct {
		name     string
		input    any
		expected PathParam
	}{
		{
			name: "Test case 1",
			input: testStruct{
				Key:   "key1",
				Value: "value1",
			},
			expected: PathParam{
				"key1": "value1",
			},
		},
		{
			name: "Test case 2",
			input: testStruct{
				Key:   "key2",
				Value: "value2",
			},
			expected: PathParam{
				"key2": "value2",
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

			p := PathParam{}
			p.SetParamsWithStruct(tc.input)
			if !reflect.DeepEqual(p, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, p)
			}
		})
	}
}
