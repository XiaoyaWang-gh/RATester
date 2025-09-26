package hugo

import (
	"fmt"
	"testing"
)

func TestReleaseVersion_1(t *testing.T) {
	type testCase struct {
		name     string
		input    Version
		expected Version
	}

	testCases := []testCase{
		{
			name:     "Test with version 1.2.3",
			input:    Version{1, 2, 3, "+dev"},
			expected: Version{1, 2, 3, ""},
		},
		{
			name:     "Test with version 0.0.0",
			input:    Version{0, 0, 0, "+dev"},
			expected: Version{0, 0, 0, ""},
		},
		{
			name:     "Test with version 4.5.6",
			input:    Version{4, 5, 6, "+dev"},
			expected: Version{4, 5, 6, ""},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.input.ReleaseVersion()
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
