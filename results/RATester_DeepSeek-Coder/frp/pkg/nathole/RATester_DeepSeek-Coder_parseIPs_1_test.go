package nathole

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseIPs_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "Test with valid IP addresses",
			input:    []string{"192.168.1.1:8080", "10.0.0.1:80", "172.16.0.1:22"},
			expected: []string{"192.168.1.1", "10.0.0.1", "172.16.0.1"},
		},
		{
			name:     "Test with invalid IP addresses",
			input:    []string{"192.168.1.1", "10.0.0.1:80", "172.16.0.1:22"},
			expected: []string{"192.168.1.1", "172.16.0.1"},
		},
		{
			name:     "Test with empty input",
			input:    []string{},
			expected: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := parseIPs(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
