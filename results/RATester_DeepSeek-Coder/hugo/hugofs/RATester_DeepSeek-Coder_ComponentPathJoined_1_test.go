package hugofs

import (
	"fmt"
	"testing"
)

func TestComponentPathJoined_1(t *testing.T) {
	t.Run("TestComponentPathJoined", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testCases := []struct {
			name     string
			input    ComponentPath
			expected string
		}{
			{
				name: "Test case 1",
				input: ComponentPath{
					Component: "component1",
					Path:      "path1",
				},
				expected: "component1/path1",
			},
			{
				name: "Test case 2",
				input: ComponentPath{
					Component: "component2",
					Path:      "path2",
				},
				expected: "component2/path2",
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {

				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered in main", r)
					}
				}()

				result := tc.input.ComponentPathJoined()
				if result != tc.expected {
					t.Errorf("Expected %s, got %s", tc.expected, result)
				}
			})
		}
	})
}
