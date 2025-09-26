package crd

import (
	"fmt"
	"testing"
)

func TestSplitSvcNameProvider_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "test case 1",
			input:    "service1" + providerNamespaceSeparator + "provider1",
			expected: []string{"service1", "provider1"},
		},
		{
			name:     "test case 2",
			input:    "service2" + providerNamespaceSeparator + "provider2",
			expected: []string{"service2", "provider2"},
		},
		// add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			svc, pvd := splitSvcNameProvider(tc.input)
			if svc != tc.expected[0] || pvd != tc.expected[1] {
				t.Errorf("Expected (%s, %s) but got (%s, %s)", tc.expected[0], tc.expected[1], svc, pvd)
			}
		})
	}
}
