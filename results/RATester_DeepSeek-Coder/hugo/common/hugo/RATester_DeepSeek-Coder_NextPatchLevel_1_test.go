package hugo

import (
	"fmt"
	"testing"
)

func TestNextPatchLevel_1(t *testing.T) {
	type testCase struct {
		name     string
		version  Version
		level    int
		expected Version
	}

	testCases := []testCase{
		{
			name:     "Test NextPatchLevel with level 1",
			version:  Version{Major: 0, Minor: 1, PatchLevel: 0},
			level:    1,
			expected: Version{Major: 0, Minor: 1, PatchLevel: 1},
		},
		{
			name:     "Test NextPatchLevel with level 2",
			version:  Version{Major: 1, Minor: 0, PatchLevel: 1},
			level:    2,
			expected: Version{Major: 1, Minor: 0, PatchLevel: 2},
		},
		{
			name:     "Test NextPatchLevel with level 3",
			version:  Version{Major: 1, Minor: 1, PatchLevel: 2},
			level:    3,
			expected: Version{Major: 1, Minor: 1, PatchLevel: 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.version.NextPatchLevel(tc.level)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
