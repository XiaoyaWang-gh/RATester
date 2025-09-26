package gin

import (
	"fmt"
	"testing"
)

func TestGetUint64_1(t *testing.T) {
	type test struct {
		name     string
		key      string
		expected uint64
		context  *Context
	}

	tests := []test{
		{
			name:     "Test case 1",
			key:      "key1",
			expected: 1,
			context: &Context{
				Keys: map[string]any{
					"key1": uint64(1),
				},
			},
		},
		{
			name:     "Test case 2",
			key:      "key2",
			expected: 2,
			context: &Context{
				Keys: map[string]any{
					"key2": uint64(2),
				},
			},
		},
		{
			name:     "Test case 3",
			key:      "key3",
			expected: 3,
			context: &Context{
				Keys: map[string]any{
					"key3": uint64(3),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tt.context.GetUint64(tt.key)
			if got != tt.expected {
				t.Errorf("Expected %d, but got %d", tt.expected, got)
			}
		})
	}
}
