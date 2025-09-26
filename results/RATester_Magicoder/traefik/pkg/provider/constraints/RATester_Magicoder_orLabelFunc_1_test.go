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
			name: "both true",
			a: func(labels map[string]string) bool {
				return true
			},
			b: func(labels map[string]string) bool {
				return true
			},
			labels:   map[string]string{"key": "value"},
			expected: true,
		},
		{
			name: "a true, b false",
			a: func(labels map[string]string) bool {
				return true
			},
			b: func(labels map[string]string) bool {
				return false
			},
			labels:   map[string]string{"key": "value"},
			expected: true,
		},
		{
			name: "a false, b true",
			a: func(labels map[string]string) bool {
				return false
			},
			b: func(labels map[string]string) bool {
				return true
			},
			labels:   map[string]string{"key": "value"},
			expected: true,
		},
		{
			name: "both false",
			a: func(labels map[string]string) bool {
				return false
			},
			b: func(labels map[string]string) bool {
				return false
			},
			labels:   map[string]string{"key": "value"},
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
