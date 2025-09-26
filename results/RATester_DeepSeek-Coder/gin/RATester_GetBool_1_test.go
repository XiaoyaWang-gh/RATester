package gin

import (
	"fmt"
	"testing"
)

func TestGetBool_1(t *testing.T) {
	type testCase struct {
		name     string
		context  *Context
		key      string
		expected bool
	}

	testCases := []testCase{
		{
			name: "TestGetBool_ExistingKey",
			context: &Context{
				Keys: map[string]any{
					"key1": true,
				},
			},
			key:      "key1",
			expected: true,
		},
		{
			name: "TestGetBool_NonExistingKey",
			context: &Context{
				Keys: map[string]any{
					"key2": false,
				},
			},
			key:      "key3",
			expected: false,
		},
		{
			name: "TestGetBool_NilValue",
			context: &Context{
				Keys: map[string]any{
					"key4": nil,
				},
			},
			key:      "key4",
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

			result := tc.context.GetBool(tc.key)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
