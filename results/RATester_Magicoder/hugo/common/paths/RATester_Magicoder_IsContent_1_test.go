package paths

import (
	"fmt"
	"testing"
)

func TestIsContent_1(t *testing.T) {
	type test struct {
		name     string
		path     *Path
		expected bool
	}

	tests := []test{
		{
			name: "Test case 1",
			path: &Path{
				bundleType: PathTypeContentResource,
			},
			expected: true,
		},
		{
			name: "Test case 2",
			path: &Path{
				bundleType: PathTypeContentData,
			},
			expected: true,
		},
		{
			name: "Test case 3",
			path: &Path{
				bundleType: PathTypeContentResource + 1,
			},
			expected: false,
		},
		{
			name: "Test case 4",
			path: &Path{
				bundleType: PathTypeContentData + 1,
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

			result := tc.path.IsContent()
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
