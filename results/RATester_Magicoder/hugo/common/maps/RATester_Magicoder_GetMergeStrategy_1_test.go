package maps

import (
	"fmt"
	"testing"
)

func TestGetMergeStrategy_1(t *testing.T) {
	testCases := []struct {
		name     string
		params   Params
		expected ParamsMergeStrategy
		found    bool
	}{
		{
			name: "MergeStrategyKey exists",
			params: Params{
				MergeStrategyKey: ParamsMergeStrategyShallow,
			},
			expected: ParamsMergeStrategyShallow,
			found:    true,
		},
		{
			name: "MergeStrategyKey does not exist",
			params: Params{
				"otherKey": "otherValue",
			},
			expected: ParamsMergeStrategyShallow,
			found:    false,
		},
		{
			name: "MergeStrategyKey is not a ParamsMergeStrategy",
			params: Params{
				MergeStrategyKey: "not a ParamsMergeStrategy",
			},
			expected: ParamsMergeStrategyShallow,
			found:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			actual, found := tc.params.GetMergeStrategy()
			if actual != tc.expected || found != tc.found {
				t.Errorf("Expected (%v, %v), but got (%v, %v)", tc.expected, tc.found, actual, found)
			}
		})
	}
}
