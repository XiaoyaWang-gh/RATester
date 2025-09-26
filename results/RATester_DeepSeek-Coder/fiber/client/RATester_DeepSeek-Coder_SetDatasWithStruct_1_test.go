package client

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestSetDatasWithStruct_1(t *testing.T) {
	type TestStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	testCases := []struct {
		name     string
		input    any
		expected map[string]string
	}{
		{
			name: "Test with valid struct",
			input: TestStruct{
				Name: "John Doe",
				Age:  30,
			},
			expected: map[string]string{
				"name": "John Doe",
				"age":  "30",
			},
		},
		{
			name:  "Test with nil",
			input: nil,
			expected: map[string]string{
				"name": "",
				"age":  "",
			},
		},
		{
			name: "Test with empty struct",
			input: TestStruct{
				Name: "",
				Age:  0,
			},
			expected: map[string]string{
				"name": "",
				"age":  "0",
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

			f := &FormData{Args: &fasthttp.Args{}}
			f.SetDatasWithStruct(tc.input)

			actual := make(map[string]string)
			f.VisitAll(func(key []byte, value []byte) {
				actual[string(key)] = string(value)
			})

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
