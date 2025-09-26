package types

import (
	"fmt"
	"testing"
)

func TestToString_5(t *testing.T) {
	t.Run("TestToString", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		type testCase struct {
			name     string
			input    any
			expected string
		}

		testCases := []testCase{
			{
				name:     "Test case 1: string input",
				input:    "Hello, world",
				expected: "Hello, world",
			},
			{
				name:     "Test case 2: int input",
				input:    123,
				expected: "123",
			},
			{
				name:     "Test case 3: float input",
				input:    123.456,
				expected: "123.456",
			},
			{
				name:     "Test case 4: bool input",
				input:    true,
				expected: "true",
			},
			{
				name:     "Test case 5: nil input",
				input:    nil,
				expected: "<nil>",
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {

				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered in main", r)
					}
				}()

				result := ToString(tc.input)
				if result != tc.expected {
					t.Errorf("Expected %s, but got %s", tc.expected, result)
				}
			})
		}
	})
}
