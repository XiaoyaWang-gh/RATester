package crd

import (
	"fmt"
	"testing"
)

func TestSplitSvcNameProvider_1(t *testing.T) {
	type test struct {
		name     string
		input    string
		expected []string
	}

	tests := []test{
		{
			name:     "Test case 1",
			input:    "service1.provider1",
			expected: []string{"service1", "provider1"},
		},
		{
			name:     "Test case 2",
			input:    "service2.provider2",
			expected: []string{"service2", "provider2"},
		},
		{
			name:     "Test case 3",
			input:    "service3.provider3",
			expected: []string{"service3", "provider3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			svc, pvd := splitSvcNameProvider(tt.input)
			if svc != tt.expected[0] || pvd != tt.expected[1] {
				t.Errorf("Expected (%s, %s), got (%s, %s)", tt.expected[0], tt.expected[1], svc, pvd)
			}
		})
	}
}
