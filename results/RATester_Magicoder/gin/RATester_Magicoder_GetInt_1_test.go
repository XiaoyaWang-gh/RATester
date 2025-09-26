package gin

import (
	"fmt"
	"testing"
)

func TestGetInt_1(t *testing.T) {
	type testCase struct {
		name     string
		context  *Context
		key      string
		expected int
	}

	testCases := []testCase{
		{
			name: "Get existing int",
			context: &Context{
				Keys: map[string]any{
					"key1": 123,
				},
			},
			key:      "key1",
			expected: 123,
		},
		{
			name: "Get non-existing int",
			context: &Context{
				Keys: map[string]any{
					"key2": "not an int",
				},
			},
			key:      "key2",
			expected: 0,
		},
		{
			name: "Get non-existing key",
			context: &Context{
				Keys: map[string]any{
					"key3": 456,
				},
			},
			key:      "key4",
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.context.GetInt(tc.key)
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}
