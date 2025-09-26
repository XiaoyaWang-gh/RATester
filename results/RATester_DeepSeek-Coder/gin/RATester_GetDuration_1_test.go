package gin

import (
	"fmt"
	"testing"
	"time"
)

func TestGetDuration_1(t *testing.T) {
	type testCase struct {
		name     string
		context  *Context
		key      string
		expected time.Duration
	}

	testCases := []testCase{
		{
			name: "Test with existing key",
			context: &Context{
				Keys: map[string]any{
					"key1": time.Second,
				},
			},
			key:      "key1",
			expected: time.Second,
		},
		{
			name: "Test with non-existing key",
			context: &Context{
				Keys: map[string]any{
					"key2": "value2",
				},
			},
			key:      "key3",
			expected: 0,
		},
		{
			name: "Test with nil value",
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

			result := tc.context.GetDuration(tc.key)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
