package echo

import (
	"fmt"
	"testing"
)

func TestBindData_1(t *testing.T) {
	b := &DefaultBinder{}

	type TestStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	testCases := []struct {
		name        string
		destination interface{}
		data        map[string][]string
		tag         string
		expected    TestStruct
	}{
		{
			name:        "Test case 1",
			destination: &TestStruct{},
			data: map[string][]string{
				"name": {"John"},
				"age":  {"30"},
			},
			tag: "json",
			expected: TestStruct{
				Name: "John",
				Age:  30,
			},
		},
		{
			name:        "Test case 2",
			destination: &TestStruct{},
			data: map[string][]string{
				"name": {"Jane"},
				"age":  {"25"},
			},
			tag: "json",
			expected: TestStruct{
				Name: "Jane",
				Age:  25,
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

			err := b.bindData(tc.destination, tc.data, tc.tag)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			result := *tc.destination.(*TestStruct)
			if result != tc.expected {
				t.Errorf("Expected: %v, got: %v", tc.expected, result)
			}
		})
	}
}
