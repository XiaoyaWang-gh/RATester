package hugolib

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestPlainWords_1(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		name     string
		input    *pageContentOutput
		expected []string
	}{
		{
			name:  "Test case 1",
			input: &pageContentOutput{
				// Initialize your input here
			},
			expected: []string{
				// Expected output here
			},
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

			result := tc.input.PlainWords(ctx)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got: %v", tc.expected, result)
			}
		})
	}
}
