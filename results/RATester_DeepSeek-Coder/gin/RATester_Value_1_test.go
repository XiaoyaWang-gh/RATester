package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestValue_1(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		context  *Context
		key      any
		expected any
	}

	testCases := []testCase{
		{
			name: "Test with ContextRequestKey",
			context: &Context{
				Request: &http.Request{},
			},
			key:      ContextRequestKey,
			expected: &http.Request{},
		},
		{
			name: "Test with ContextKey",
			context: &Context{
				Request: &http.Request{},
			},
			key:      ContextKey,
			expected: &Context{},
		},
		{
			name: "Test with string key",
			context: &Context{
				Keys: map[string]any{
					"key": "value",
				},
			},
			key:      "key",
			expected: "value",
		},
		{
			name: "Test with unknown key",
			context: &Context{
				Keys: map[string]any{
					"key": "value",
				},
			},
			key:      "unknown",
			expected: nil,
		},
		{
			name: "Test with no request context",
			context: &Context{
				Request: &http.Request{},
			},
			key:      "key",
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.context.Value(tc.key)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
