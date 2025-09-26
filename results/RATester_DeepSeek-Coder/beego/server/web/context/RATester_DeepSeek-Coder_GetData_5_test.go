package context

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetData_5(t *testing.T) {
	type testStruct struct {
		Name string
		Age  int
	}

	testCases := []struct {
		name     string
		input    *BeegoInput
		key      interface{}
		expected interface{}
	}{
		{
			name: "Test case 1: key exists",
			input: &BeegoInput{
				data: map[interface{}]interface{}{
					"key1": "value1",
					"key2": 123,
					"key3": testStruct{
						Name: "test",
						Age:  30,
					},
				},
			},
			key:      "key2",
			expected: 123,
		},
		{
			name: "Test case 2: key does not exist",
			input: &BeegoInput{
				data: map[interface{}]interface{}{
					"key1": "value1",
					"key2": 123,
					"key3": testStruct{
						Name: "test",
						Age:  30,
					},
				},
			},
			key:      "key4",
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.input.GetData(tc.key)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
