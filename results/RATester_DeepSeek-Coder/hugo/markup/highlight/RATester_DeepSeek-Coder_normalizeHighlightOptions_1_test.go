package highlight

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNormalizeHighlightOptions_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    map[string]any
		expected map[string]any
	}{
		{
			name:  "nil input",
			input: nil,
			expected: map[string]any{
				"linosStartKey": 1,
			},
		},
		{
			name: "normal case",
			input: map[string]any{
				"linosStartKey": 2,
				"noHl":          "true",
				"lineNos":       "false",
				"hlLines":       [][2]int{{3, 4}},
			},
			expected: map[string]any{
				"linosstartkey":  2,
				"noHl":           true,
				"lineNos":        false,
				"hlLines_parsed": [][2]int{{4, 5}},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			normalizeHighlightOptions(tc.input)
			if !reflect.DeepEqual(tc.input, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, tc.input)
			}
		})
	}
}
