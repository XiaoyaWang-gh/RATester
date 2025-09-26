package gin

import (
	"fmt"
	"testing"
)

func TestGetInt_1(t *testing.T) {
	type testCase struct {
		name     string
		key      string
		expected int
		context  *Context
	}

	testCases := []testCase{
		{
			name:     "Test case 1",
			key:      "key1",
			expected: 1,
			context: &Context{
				Keys: map[string]any{
					"key1": 1,
				},
			},
		},
		{
			name:     "Test case 2",
			key:      "key2",
			expected: 0,
			context: &Context{
				Keys: map[string]any{
					"key2": "value2",
				},
			},
		},
		{
			name:     "Test case 3",
			key:      "key3",
			expected: 0,
			context: &Context{
				Keys: map[string]any{
					"key3": nil,
				},
			},
		},
		{
			name:     "Test case 4",
			key:      "key4",
			expected: 0,
			context: &Context{
				Keys: map[string]any{
					"key4": 3.14,
				},
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

			result := tc.context.GetInt(tc.key)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
