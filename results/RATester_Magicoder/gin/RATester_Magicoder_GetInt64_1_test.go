package gin

import (
	"fmt"
	"testing"
)

func TestGetInt64_1(t *testing.T) {
	type testCase struct {
		name     string
		context  *Context
		key      string
		expected int64
	}

	testCases := []testCase{
		{
			name: "Get existing int64",
			context: &Context{
				Keys: map[string]any{
					"key1": int64(123),
				},
			},
			key:      "key1",
			expected: 123,
		},
		{
			name: "Get non-existing int64",
			context: &Context{
				Keys: map[string]any{
					"key2": "not an int64",
				},
			},
			key:      "key2",
			expected: 0,
		},
		{
			name: "Get int64 from empty context",
			context: &Context{
				Keys: map[string]any{},
			},
			key:      "key3",
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

			result := tc.context.GetInt64(tc.key)
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}
