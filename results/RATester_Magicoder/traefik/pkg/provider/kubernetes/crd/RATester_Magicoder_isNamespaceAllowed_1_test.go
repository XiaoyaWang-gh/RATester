package crd

import (
	"fmt"
	"testing"
)

func TestIsNamespaceAllowed_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type testCase struct {
		allowCrossNamespace bool
		parentNamespace     string
		namespace           string
		expected            bool
	}

	testCases := []testCase{
		{
			allowCrossNamespace: true,
			parentNamespace:     "parent",
			namespace:           "child",
			expected:            true,
		},
		{
			allowCrossNamespace: false,
			parentNamespace:     "parent",
			namespace:           "child",
			expected:            false,
		},
		{
			allowCrossNamespace: true,
			parentNamespace:     "parent",
			namespace:           "parent",
			expected:            true,
		},
		{
			allowCrossNamespace: false,
			parentNamespace:     "parent",
			namespace:           "parent",
			expected:            true,
		},
	}

	for _, tc := range testCases {
		result := isNamespaceAllowed(tc.allowCrossNamespace, tc.parentNamespace, tc.namespace)
		if result != tc.expected {
			t.Errorf("isNamespaceAllowed(%v, %v, %v) = %v; want %v", tc.allowCrossNamespace, tc.parentNamespace, tc.namespace, result, tc.expected)
		}
	}
}
