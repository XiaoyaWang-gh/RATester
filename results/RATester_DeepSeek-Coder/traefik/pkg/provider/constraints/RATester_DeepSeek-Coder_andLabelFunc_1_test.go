package constraints

import (
	"fmt"
	"testing"
)

func TestAndLabelFunc_1(t *testing.T) {
	testCases := []struct {
		name     string
		a        constraintLabelFunc
		b        constraintLabelFunc
		labels   map[string]string
		expected bool
	}{
		{
			name: "Test case 1",
			a: func(labels map[string]string) bool {
				_, ok := labels["key1"]
				return ok
			},
			b: func(labels map[string]string) bool {
				_, ok := labels["key2"]
				return ok
			},
			labels:   map[string]string{"key1": "value1"},
			expected: false,
		},
		{
			name: "Test case 2",
			a: func(labels map[string]string) bool {
				_, ok := labels["key1"]
				return ok
			},
			b: func(labels map[string]string) bool {
				_, ok := labels["key1"]
				return ok
			},
			labels:   map[string]string{"key1": "value1"},
			expected: true,
		},
		{
			name: "Test case 3",
			a: func(labels map[string]string) bool {
				_, ok := labels["key1"]
				return ok
			},
			b: func(labels map[string]string) bool {
				_, ok := labels["key1"]
				return ok
			},
			labels:   map[string]string{"key2": "value2"},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := andLabelFunc(tc.a, tc.b)(tc.labels)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
