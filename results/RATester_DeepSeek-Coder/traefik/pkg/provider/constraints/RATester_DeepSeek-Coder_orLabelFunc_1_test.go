package constraints

import (
	"fmt"
	"testing"
)

func TestOrLabelFunc_1(t *testing.T) {
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
				return labels["field name"] == "value1"
			},
			b: func(labels map[string]string) bool {
				return labels["field name"] == "value2"
			},
			labels: map[string]string{
				"field name": "value1",
			},
			expected: true,
		},
		{
			name: "Test case 2",
			a: func(labels map[string]string) bool {
				return labels["field name"] == "value1"
			},
			b: func(labels map[string]string) bool {
				return labels["field name"] == "value2"
			},
			labels: map[string]string{
				"field name": "value2",
			},
			expected: true,
		},
		{
			name: "Test case 3",
			a: func(labels map[string]string) bool {
				return labels["field name"] == "value1"
			},
			b: func(labels map[string]string) bool {
				return labels["field name"] == "value2"
			},
			labels: map[string]string{
				"field name": "value3",
			},
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

			result := orLabelFunc(tc.a, tc.b)(tc.labels)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
