package glob

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestNewFilenameFilterForInclusionFunc_1(t *testing.T) {
	testCases := []struct {
		name          string
		shouldInclude func(filename string) bool
		expected      *FilenameFilter
	}{
		{
			name: "Test case 1",
			shouldInclude: func(filename string) bool {
				return strings.HasSuffix(filename, ".go")
			},
			expected: &FilenameFilter{
				shouldInclude: func(filename string) bool {
					return strings.HasSuffix(filename, ".go")
				},
				isWindows: isWindows,
			},
		},
		{
			name: "Test case 2",
			shouldInclude: func(filename string) bool {
				return strings.HasSuffix(filename, ".txt")
			},
			expected: &FilenameFilter{
				shouldInclude: func(filename string) bool {
					return strings.HasSuffix(filename, ".txt")
				},
				isWindows: isWindows,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := NewFilenameFilterForInclusionFunc(tc.shouldInclude)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
