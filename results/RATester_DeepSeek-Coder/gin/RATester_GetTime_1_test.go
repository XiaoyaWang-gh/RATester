package gin

import (
	"fmt"
	"testing"
	"time"
)

func TestGetTime_1(t *testing.T) {
	type testCase struct {
		name     string
		context  *Context
		key      string
		expected time.Time
	}

	testCases := []testCase{
		{
			name: "Get existing time",
			context: &Context{
				Keys: map[string]any{
					"key1": time.Now(),
				},
			},
			key:      "key1",
			expected: time.Now(),
		},
		{
			name: "Get non-existing time",
			context: &Context{
				Keys: map[string]any{
					"key2": "not a time",
				},
			},
			key:      "key2",
			expected: time.Time{},
		},
		{
			name: "Get time from empty context",
			context: &Context{
				Keys: map[string]any{},
			},
			key:      "key3",
			expected: time.Time{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.context.GetTime(tc.key)
			if !result.Equal(tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
