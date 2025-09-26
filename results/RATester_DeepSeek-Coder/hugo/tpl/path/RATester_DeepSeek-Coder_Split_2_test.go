package path

import (
	"errors"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/paths"
)

func TestSplit_2(t *testing.T) {
	type testCase struct {
		name     string
		path     any
		expected paths.DirFile
		err      error
	}

	tests := []testCase{
		{
			name:     "valid path",
			path:     "/path/to/file",
			expected: paths.DirFile{Dir: "/path/to/", File: "file"},
			err:      nil,
		},
		{
			name:     "invalid path",
			path:     123,
			expected: paths.DirFile{},
			err:      errors.New("cast: unable to cast 123 of type int to string"),
		},
		{
			name:     "empty path",
			path:     "",
			expected: paths.DirFile{},
			err:      nil,
		},
	}

	ns := &Namespace{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := ns.Split(tc.path)
			if err != nil && tc.err == nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if err != nil && tc.err != nil && err.Error() != tc.err.Error() {
				t.Errorf("Expected error %v, got %v", tc.err, err)
			}
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
