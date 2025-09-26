package helloworld

import (
	fmt "fmt"
	"testing"
)

func TestGetMessage_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    *HelloReply
		expected string
	}{
		{
			name:     "Test case 1: Happy path",
			input:    &HelloReply{Message: "Hello, World!"},
			expected: "Hello, World!",
		},
		{
			name:     "Test case 2: Nil input",
			input:    nil,
			expected: "",
		},
		{
			name:     "Test case 3: Empty message",
			input:    &HelloReply{Message: ""},
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.input.GetMessage()
			if result != tc.expected {
				t.Errorf("Expected: %s, but got: %s", tc.expected, result)
			}
		})
	}
}
