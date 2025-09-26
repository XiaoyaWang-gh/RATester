package gin

import (
	"fmt"
	"testing"
)

func TestGetUint_1(t *testing.T) {
	type testCase struct {
		name     string
		context  *Context
		key      string
		expected uint
	}

	testCases := []testCase{
		{
			name: "Test case 1",
			context: &Context{
				Keys: map[string]any{
					"key1": uint(1),
				},
			},
			key:      "key1",
			expected: 1,
		},
		{
			name: "Test case 2",
			context: &Context{
				Keys: map[string]any{
					"key2": uint(2),
				},
			},
			key:      "key2",
			expected: 2,
		},
		{
			name: "Test case 3",
			context: &Context{
				Keys: map[string]any{
					"key3": "not a uint",
				},
			},
			key:      "key3",
			expected: 0,
		},
		{
			name: "Test case 4",
			context: &Context{
				Keys: map[string]any{
					"key4": nil,
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

			result := tc.context.GetUint(tc.key)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
