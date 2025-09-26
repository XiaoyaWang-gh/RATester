package gin

import (
	"fmt"
	"testing"
)

func TestGetInt64_1(t *testing.T) {
	type testCase struct {
		name     string
		key      string
		expected int64
		context  *Context
	}

	testCases := []testCase{
		{
			name:     "Test with existing key",
			key:      "existingKey",
			expected: 123,
			context: &Context{
				Keys: map[string]any{
					"existingKey": int64(123),
				},
			},
		},
		{
			name:     "Test with non-existing key",
			key:      "nonExistingKey",
			expected: 0,
			context: &Context{
				Keys: map[string]any{
					"existingKey": int64(123),
				},
			},
		},
		{
			name:     "Test with nil value",
			key:      "nilKey",
			expected: 0,
			context: &Context{
				Keys: map[string]any{
					"nilKey": nil,
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

			result := tc.context.GetInt64(tc.key)
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}
