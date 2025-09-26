package path

import (
	"errors"
	"fmt"
	"testing"
)

func TestDir_2(t *testing.T) {
	type testCase struct {
		name     string
		path     any
		expected string
		err      error
	}

	testCases := []testCase{
		{
			name:     "valid path",
			path:     "/foo/bar/baz.txt",
			expected: "/foo/bar",
			err:      nil,
		},
		{
			name:     "invalid path",
			path:     123,
			expected: "",
			err:      errors.New("cast: unable to cast 123 of type int to string"),
		},
	}

	ns := &Namespace{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			actual, err := ns.Dir(tc.path)
			if actual != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, actual)
			}
			if err != nil && err.Error() != tc.err.Error() {
				t.Errorf("expected error %q, got %q", tc.err, err)
			}
		})
	}
}
