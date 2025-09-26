package maps

import (
	"fmt"
	"testing"
)

func TestDeleteMergeStrategy_1(t *testing.T) {
	testCases := []struct {
		name     string
		params   Params
		expected bool
	}{
		{
			name:     "delete existing key",
			params:   Params{MergeStrategyKey: "existing"},
			expected: true,
		},
		{
			name:     "delete non-existing key",
			params:   Params{},
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
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
