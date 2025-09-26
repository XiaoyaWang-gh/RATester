package web

import (
	"fmt"
	"testing"
)

func TestWithResetParams_1(t *testing.T) {
	t.Run("TestWithResetParams", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testCases := []struct {
			name     string
			input    bool
			expected bool
		}{
			{
				name:     "Test case 1",
				input:    true,
				expected: true,
			},
			{
				name:     "Test case 2",
				input:    false,
				expected: false,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {

				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered in main", r)
					}
				}()

				opts := &filterOpts{}
				WithResetParams(tc.input)(opts)
				if opts.resetParams != tc.expected {
					t.Errorf("Expected %v, got %v", tc.expected, opts.resetParams)
				}
			})
		}
	})
}
