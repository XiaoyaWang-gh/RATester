package maps

import (
	"fmt"
	"testing"
)

func TestSetMergeStrategy_1(t *testing.T) {
	testCases := []struct {
		name     string
		strategy ParamsMergeStrategy
		expected error
	}{
		{
			name:     "valid strategy",
			strategy: ParamsMergeStrategyDeep,
			expected: nil,
		},
		{
			name:     "invalid strategy",
			strategy: "invalid",
			expected: fmt.Errorf("invalid merge strategy %q", "invalid"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := Params{}
			p.SetMergeStrategy(tc.strategy)

			if tc.expected != nil {
				if _, ok := p[MergeStrategyKey]; ok {
					t.Errorf("Expected error, but strategy was set")
				}
			} else {
				if _, ok := p[MergeStrategyKey]; !ok {
					t.Errorf("Expected strategy, but none was set")
				}
			}
		})
	}
}
