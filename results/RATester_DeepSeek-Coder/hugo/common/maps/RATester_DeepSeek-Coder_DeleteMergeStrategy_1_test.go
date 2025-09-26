package maps

import (
	"fmt"
	"testing"
)

func TestDeleteMergeStrategy_1(t *testing.T) {
	type testCase struct {
		name     string
		params   Params
		expected bool
	}

	testCases := []testCase{
		{
			name: "Delete existing merge strategy",
			params: Params{
				MergeStrategyKey: "merge",
				"other":          "value",
			},
			expected: true,
		},
		{
			name: "Delete non-existing merge strategy",
			params: Params{
				"other": "value",
			},
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

			result := tc.params.DeleteMergeStrategy()
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
