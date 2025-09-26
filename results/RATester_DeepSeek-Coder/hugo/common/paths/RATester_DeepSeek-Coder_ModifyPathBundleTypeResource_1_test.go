package paths

import (
	"fmt"
	"testing"
)

func TestModifyPathBundleTypeResource_1(t *testing.T) {
	type testCase struct {
		name     string
		path     *Path
		expected PathType
	}

	testCases := []testCase{
		{
			name: "Content path",
			path: &Path{
				s:          "content/test.md",
				bundleType: PathTypeFile,
			},
			expected: PathTypeContentResource,
		},
		{
			name: "Non-content path",
			path: &Path{
				s:          "static/test.txt",
				bundleType: PathTypeFile,
			},
			expected: PathTypeFile,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ModifyPathBundleTypeResource(tc.path)
			if tc.path.bundleType != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, tc.path.bundleType)
			}
		})
	}
}
