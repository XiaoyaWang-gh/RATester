package resources

import (
	"fmt"
	"testing"
)

func TestIsContentChanged_1(t *testing.T) {
	type testCase struct {
		name     string
		update   *transformationUpdate
		expected bool
	}

	testCases := []testCase{
		{
			name: "Both content and sourceFilename are nil",
			update: &transformationUpdate{
				content:        nil,
				sourceFilename: nil,
			},
			expected: false,
		},
		{
			name: "Only content is nil",
			update: &transformationUpdate{
				content:        nil,
				sourceFilename: &[]string{"test"}[0],
			},
			expected: true,
		},
		{
			name: "Only sourceFilename is nil",
			update: &transformationUpdate{
				content:        &[]string{"test"}[0],
				sourceFilename: nil,
			},
			expected: true,
		},
		{
			name: "Both content and sourceFilename are not nil",
			update: &transformationUpdate{
				content:        &[]string{"test"}[0],
				sourceFilename: &[]string{"test"}[0],
			},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.update.isContentChanged()
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
