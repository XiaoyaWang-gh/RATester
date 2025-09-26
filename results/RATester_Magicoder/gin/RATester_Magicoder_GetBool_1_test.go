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
			name: "GetBool_True",
			context: &Context{
				Keys: map[string]any{
					"key1": true,
				},
			},
			key:      "key1",
			expected: true,
		},
		{
			name: "GetBool_False",
			context: &Context{
				Keys: map[string]any{
					"key2": false,
				},
			},
			key:      "key2",
			expected: false,
		},
		{
			name: "GetBool_NotExist",
			context: &Context{
				Keys: map[string]any{
					"key3": "value3",
				},
			},
			key:      "key4",
			expected: false,
		},
		{
			name: "GetBool_WrongType",
			context: &Context{
				Keys: map[string]any{
					"key5": 123,
				},
			},
			key:      "key5",
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
