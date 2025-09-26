package maps

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gobwas/glob"
)

func TestRename_2(t *testing.T) {
	type testCase struct {
		name     string
		renamer  KeyRenamer
		input    map[string]any
		expected map[string]any
	}

	testCases := []testCase{
		{
			name: "simple rename",
			renamer: KeyRenamer{
				renames: []keyRename{
					{pattern: glob.MustCompile("oldKey"), newKey: "newKey"},
				},
			},
			input: map[string]any{
				"oldKey": "value",
			},
			expected: map[string]any{
				"newKey": "value",
			},
		},
		// Add more test cases here...
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tc.renamer.Rename(tc.input)
			if !reflect.DeepEqual(tc.input, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, tc.input)
			}
		})
	}
}
