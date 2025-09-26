package deployconfig

import (
	"errors"
	"fmt"
	"testing"
)

func TestParseIncludeExclude_1(t *testing.T) {
	type testCase struct {
		name     string
		target   *Target
		expected error
	}

	testCases := []testCase{
		{
			name: "Include is valid",
			target: &Target{
				Include: "*",
			},
			expected: nil,
		},
		{
			name: "Exclude is valid",
			target: &Target{
				Exclude: "*",
			},
			expected: nil,
		},
		{
			name: "Include is invalid",
			target: &Target{
				Include: "**",
			},
			expected: fmt.Errorf("invalid deployment.target.include %q: %v", "**", errors.New("unmatched '*'")),
		},
		{
			name: "Exclude is invalid",
			target: &Target{
				Exclude: "**",
			},
			expected: fmt.Errorf("invalid deployment.target.exclude %q: %v", "**", errors.New("unmatched '*'")),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := tc.target.ParseIncludeExclude()
			if err != nil && tc.expected != nil {
				if err.Error() != tc.expected.Error() {
					t.Errorf("Expected error %v, got %v", tc.expected, err)
				}
			} else if err != tc.expected {
				t.Errorf("Expected error %v, got %v", tc.expected, err)
			}
		})
	}
}
