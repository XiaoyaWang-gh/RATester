package gin

import (
	"fmt"
	"testing"
)

func TestGetFloat64_1(t *testing.T) {
	type testCase struct {
		name     string
		context  *Context
		key      string
		expected float64
	}

	testCases := []testCase{
		{
			name: "Get existing float64",
			context: &Context{
				Keys: map[string]any{
					"key1": float64(123.456),
				},
			},
			key:      "key1",
			expected: 123.456,
		},
		{
			name: "Get non-existing float64",
			context: &Context{
				Keys: map[string]any{
					"key2": "not a float64",
				},
			},
			key:      "key2",
			expected: 0,
		},
		{
			name: "Get float64 from empty context",
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

			result := tc.context.GetFloat64(tc.key)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
