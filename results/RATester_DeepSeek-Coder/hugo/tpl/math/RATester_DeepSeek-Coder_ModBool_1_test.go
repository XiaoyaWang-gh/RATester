package math

import (
	"errors"
	"fmt"
	"testing"
)

func TestModBool_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("TestModBool", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testCases := []struct {
			name     string
			n1       any
			n2       any
			expected bool
			err      error
		}{
			{
				name:     "Test case 1",
				n1:       10,
				n2:       3,
				expected: false,
				err:      nil,
			},
			{
				name:     "Test case 2",
				n1:       10,
				n2:       2,
				expected: true,
				err:      nil,
			},
			{
				name:     "Test case 3",
				n1:       "10",
				n2:       3,
				expected: false,
				err:      errors.New("unsupported operation: Mod"),
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {

				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered in main", r)
					}
				}()

				result, err := ns.ModBool(tc.n1, tc.n2)
				if err != nil && err.Error() != tc.err.Error() {
					t.Errorf("Expected error %v, got %v", tc.err, err)
				}
				if result != tc.expected {
					t.Errorf("Expected result %v, got %v", tc.expected, result)
				}
			})
		}
	})
}
