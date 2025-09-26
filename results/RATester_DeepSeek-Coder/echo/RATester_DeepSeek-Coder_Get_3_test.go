package echo

import (
	"fmt"
	"testing"
)

func TestGet_3(t *testing.T) {
	type testCase struct {
		name     string
		key      string
		expected interface{}
	}

	testCases := []testCase{
		{
			name:     "TestGet_ExistingKey",
			key:      "existingKey",
			expected: "existingValue",
		},
		{
			name:     "TestGet_NonExistingKey",
			key:      "nonExistingKey",
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &context{
				store: Map{
					"existingKey": "existingValue",
				},
			}

			result := c.Get(tc.key)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
