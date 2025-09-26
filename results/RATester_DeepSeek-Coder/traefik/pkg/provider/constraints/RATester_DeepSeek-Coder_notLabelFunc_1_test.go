package constraints

import (
	"fmt"
	"testing"
)

func TestNotLabelFunc_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    map[string]string
		expected bool
	}{
		{
			name:     "Test case 1",
			input:    map[string]string{"key1": "value1", "key2": "value2"},
			expected: false,
		},
		{
			name:     "Test case 2",
			input:    map[string]string{"key3": "value3", "key4": "value4"},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			fn := notLabelFunc(func(labels map[string]string) bool {
				return len(labels) > 0
			})

			result := fn(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
