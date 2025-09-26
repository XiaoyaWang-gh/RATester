package path

import (
	"errors"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/paths"
)

func TestSplit_2(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name     string
		path     any
		expected paths.DirFile
		err      error
	}{
		{
			name:     "valid path",
			path:     "/path/to/file.txt",
			expected: paths.DirFile{Dir: "/path/to/", File: "file.txt"},
			err:      nil,
		},
		{
			name:     "invalid path",
			path:     123,
			expected: paths.DirFile{},
			err:      errors.New("unsupported type for path"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := ns.Split(test.path)
			if err != nil && err.Error() != test.err.Error() {
				t.Errorf("Expected error %v, but got %v", test.err, err)
			}
			if result != test.expected {
				t.Errorf("Expected result %v, but got %v", test.expected, result)
			}
		})
	}
}
