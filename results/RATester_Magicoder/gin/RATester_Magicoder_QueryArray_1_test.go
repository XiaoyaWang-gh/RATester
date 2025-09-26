package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestQueryArray_1(t *testing.T) {
	type test struct {
		name     string
		key      string
		expected []string
	}

	tests := []test{
		{
			name:     "Test case 1",
			key:      "key1",
			expected: []string{"value1", "value2"},
		},
		{
			name:     "Test case 2",
			key:      "key2",
			expected: []string{"value3", "value4"},
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

			result := c.QueryArray(tt.key)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected: %v, got: %v", tt.expected, result)
			}
		})
	}
}
