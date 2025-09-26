package maps

import (
	"fmt"
	"testing"
)

func TestGetMergeStrategy_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		params   Params
		expected ParamsMergeStrategy
		found    bool
	}{
		{
			name:     "found",
			params:   Params{"merge_strategy": ParamsMergeStrategy("deep")},
			expected: "deep",
			found:    true,
		},
		{
			name:     "not found",
			params:   Params{"merge_strategy": "not_exist"},
			expected: "shallow",
			found:    false,
		},
		{
			name:     "not a ParamsMergeStrategy",
			params:   Params{"merge_strategy": "not_ParamsMergeStrategy"},
			expected: "shallow",
			found:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			actual, found := test.params.GetMergeStrategy()
			if actual != test.expected || found != test.found {
				t.Errorf("Expected (%v, %v), but got (%v, %v)", test.expected, test.found, actual, found)
			}
		})
	}
}
