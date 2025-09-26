package gin

import (
	"errors"
	"fmt"
	"testing"
)

func TestShouldBindUri_1(t *testing.T) {
	type testStruct struct {
		Field1 string `uri:"field1"`
		Field2 string `uri:"field2"`
	}

	testCases := []struct {
		name     string
		context  *Context
		expected error
	}{
		{
			name: "Success",
			context: &Context{
				Params: Params{
					{Key: "field1", Value: "value1"},
					{Key: "field2", Value: "value2"},
				},
			},
			expected: nil,
		},
		{
			name: "Error",
			context: &Context{
				Params: Params{
					{Key: "field1", Value: "value1"},
					{Key: "field2", Value: "value2"},
				},
			},
			expected: errors.New("some error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			obj := &testStruct{}
			err := tc.context.ShouldBindUri(obj)
			if err != tc.expected {
				t.Errorf("Expected error %v, but got %v", tc.expected, err)
			}
		})
	}
}
