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
			name: "Test case 1",
			context: &Context{
				Keys: map[string]any{
					"key1": time.Second,
				},
			},
			key:      "key1",
			expected: time.Second,
		},
		{
			name: "Test case 2",
			context: &Context{
				Keys: map[string]any{
					"key2": time.Minute,
				},
			},
			key:      "key2",
			expected: time.Minute,
		},
		{
			name: "Test case 3",
			context: &Context{
				Keys: map[string]any{
					"key3": time.Hour,
				},
			},
			key:      "key3",
			expected: time.Hour,
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
