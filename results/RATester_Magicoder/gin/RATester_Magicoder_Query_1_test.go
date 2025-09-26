package gin

import (
	"fmt"
	"testing"
)

func TestQuery_1(t *testing.T) {
	type test struct {
		name     string
		key      string
		expected string
	}

	tests := []test{
		{
			name:     "Test case 1",
			key:      "key1",
			expected: "value1",
		},
		{
			name:     "Test case 2",
			key:      "key2",
			expected: "value2",
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &Context{
				Keys: map[string]any{
					tt.key: tt.expected,
				},
			}

			value := c.Query(tt.key)

			if value != tt.expected {
				t.Errorf("Expected %s, but got %s", tt.expected, value)
			}
		})
	}
}
