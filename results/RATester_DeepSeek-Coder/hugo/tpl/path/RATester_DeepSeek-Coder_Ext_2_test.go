package path

import (
	"errors"
	"fmt"
	"testing"
)

func TestExt_2(t *testing.T) {
	type testCase struct {
		name     string
		path     any
		expected string
		err      error
	}

	testCases := []testCase{
		{
			name:     "valid path",
			path:     "/path/to/file.txt",
			expected: ".txt",
			err:      nil,
		},
		{
			name:     "invalid path",
			path:     123,
			expected: "",
			err:      errors.New("cast: unable to cast 123 of type int to string"),
		},
		{
			name:     "empty path",
			path:     "",
			expected: "",
			err:      nil,
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

			result, err := ns.Ext(tc.path)
			if tc.err != nil {
				if err == nil {
					t.Errorf("Expected error, but got nil")
				} else if err.Error() != tc.err.Error() {
					t.Errorf("Expected error %v, but got %v", tc.err, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result != tc.expected {
					t.Errorf("Expected %v, but got %v", tc.expected, result)
				}
			}
		})
	}
}
