package gin

import (
	"fmt"
	"testing"
)

func TestGetString_1(t *testing.T) {
	type testCase struct {
		name     string
		context  *Context
		key      string
		expected string
	}

	testCases := []testCase{
		{
			name: "TestGetString_ExistingKey",
			context: &Context{
				Keys: map[string]any{
					"key1": "value1",
				},
			},
			key:      "key1",
			expected: "value1",
		},
		{
			name: "TestGetString_NonExistingKey",
			context: &Context{
				Keys: map[string]any{
					"key2": "value2",
				},
			},
			key:      "key1",
			expected: "",
		},
		{
			name: "TestGetString_NilValue",
			context: &Context{
				Keys: map[string]any{
					"key3": nil,
				},
			},
			key:      "key3",
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.context.GetString(tc.key)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
