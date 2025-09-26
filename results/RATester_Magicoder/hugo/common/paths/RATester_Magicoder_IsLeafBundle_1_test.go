package paths

import (
	"fmt"
	"testing"
)

func TestIsLeafBundle_1(t *testing.T) {
	type test struct {
		name     string
		path     *Path
		expected bool
	}

	tests := []test{
		{
			name: "LeafBundle",
			path: &Path{
				bundleType: PathTypeLeaf,
			},
			expected: true,
		},
		{
			name: "NonLeafBundle",
			path: &Path{
				bundleType: 0,
			},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.path.IsLeafBundle()
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
