package protoexample

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestString_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    *Test_OptionalGroup
		expected string
	}{
		{
			name: "Test case 1",
			input: &Test_OptionalGroup{
				RequiredField: proto.String("test"),
			},
			expected: "RequiredField: \"test\"",
		},
		{
			name: "Test case 2",
			input: &Test_OptionalGroup{
				RequiredField: proto.String(""),
			},
			expected: "RequiredField: \"\"",
		},
		{
			name: "Test case 3",
			input: &Test_OptionalGroup{
				RequiredField: nil,
			},
			expected: "RequiredField: <nil>",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.input.String()
			if result != tc.expected {
				t.Errorf("Expected: %s, but got: %s", tc.expected, result)
			}
		})
	}
}
