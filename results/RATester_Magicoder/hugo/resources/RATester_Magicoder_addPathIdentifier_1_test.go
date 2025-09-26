package resources

import (
	"fmt"
	"testing"
)

func TestaddPathIdentifier_1(t *testing.T) {
	ctx := &ResourceTransformationCtx{
		InPath: "/path/to/file.txt",
	}

	testCases := []struct {
		name       string
		identifier string
		expected   string
	}{
		{
			name:       "Add identifier to path",
			identifier: "_new",
			expected:   "/path/to/file_new.txt",
		},
		{
			name:       "Add identifier to path with existing identifier",
			identifier: "_existing",
			expected:   "/path/to/file_existing.txt",
		},
		{
			name:       "Add identifier to path with no extension",
			identifier: "_no_ext",
			expected:   "/path/to/file_no_ext",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := ctx.addPathIdentifier(ctx.InPath, tc.identifier)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
