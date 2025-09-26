package path

import (
	"errors"
	"fmt"
	"testing"
)

func TestBase_2(t *testing.T) {
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
			expected: "baz.txt",
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

			result, err := ns.Base(tc.path)
			if tc.err != nil {
				if err == nil {
					t.Errorf("expected error, but got nil")
				} else if err.Error() != tc.err.Error() {
					t.Errorf("expected error %v, but got %v", tc.err, err)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result != tc.expected {
					t.Errorf("expected %v, but got %v", tc.expected, result)
				}
			}
		})
	}
}
