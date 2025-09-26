package gin

import (
	"fmt"
	"testing"
)

func TestMustGet_1(t *testing.T) {
	type test struct {
		name     string
		context  *Context
		key      string
		expected any
	}

	tests := []test{
		{
			name: "Test case 1",
			context: &Context{
				Keys: map[string]any{
					"key1": "value1",
				},
			},
			key:      "key1",
			expected: "value1",
		},
		{
			name: "Test case 2",
			context: &Context{
				Keys: map[string]any{
					"key2": "value2",
				},
			},
			key:      "key2",
			expected: "value2",
		},
		{
			name: "Test case 3",
			context: &Context{
				Keys: map[string]any{
					"key3": "value3",
				},
			},
			key:      "key3",
			expected: "value3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tt.context.MustGet(tt.key)
			if got != tt.expected {
				t.Errorf("Expected %v, but got %v", tt.expected, got)
			}
		})
	}
}
