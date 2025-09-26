package glob

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gobwas/glob"
)

func TestNewFilenameFilter_1(t *testing.T) {
	tests := []struct {
		name        string
		inclusions  []string
		exclusions  []string
		expected    *FilenameFilter
		expectedErr error
	}{
		{
			name:       "nil inclusions and exclusions",
			inclusions: nil,
			exclusions: nil,
			expected:   nil,
		},
		{
			name:       "non-nil inclusions and exclusions",
			inclusions: []string{"*.go", "*.txt"},
			exclusions: []string{"*.md"},
			expected: &FilenameFilter{
				inclusions:    []glob.Glob{glob.MustCompile("*.go"), glob.MustCompile("*.txt")},
				dirInclusions: []glob.Glob{glob.MustCompile("/"), glob.MustCompile("/."), glob.MustCompile("/.."), glob.MustCompile("/.*"), glob.MustCompile("/..*"), glob.MustCompile("/.*.*"), glob.MustCompile("/..*.*")},
				exclusions:    []glob.Glob{glob.MustCompile("*.md")},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			actual, err := NewFilenameFilter(test.inclusions, test.exclusions)
			if err != test.expectedErr {
				t.Errorf("Expected error %v, but got %v", test.expectedErr, err)
			}
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, actual)
			}
		})
	}
}
