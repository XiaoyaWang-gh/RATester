package gin

import (
	"fmt"
	"testing"
)

func TestGetUint64_1(t *testing.T) {
	type testCase struct {
		name     string
		context  *Context
		key      string
		expected uint64
	}

	testCases := []testCase{
		{
			name: "Test case 1",
			context: &Context{
				Keys: map[string]any{
					"key1": uint64(123),
				},
			},
			key:      "key1",
			expected: 123,
		},
		{
			name: "Test case 2",
			context: &Context{
				Keys: map[string]any{
					"key2": "not a uint64",
				},
			},
			key:      "key2",
			expected: 0,
		},
		{
			name: "Test case 3",
			context: &Context{
				Keys: map[string]any{
					"key3": nil,
				},
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

			result := tc.context.GetUint64(tc.key)
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}
