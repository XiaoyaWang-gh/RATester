package crypt

import (
	"fmt"
	"testing"
)

func TestGetServerName_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    *ClientHelloMsg
		expected string
	}{
		{
			name: "Test case 1",
			input: &ClientHelloMsg{
				serverName: "example.com",
			},
			expected: "example.com",
		},
		{
			name: "Test case 2",
			input: &ClientHelloMsg{
				serverName: "google.com",
			},
			expected: "google.com",
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.input.GetServerName()
			if result != tc.expected {
				t.Errorf("Expected: %s, but got: %s", tc.expected, result)
			}
		})
	}
}
