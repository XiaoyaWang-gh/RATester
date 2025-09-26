package echo

import (
	"fmt"
	"io/fs"
	"os"
	"reflect"
	"syscall"
	"testing"
)

func TestSubFS_1(t *testing.T) {
	type testCase struct {
		name     string
		fs       fs.FS
		root     string
		expected fs.FS
		err      error
	}

	testCases := []testCase{
		{
			name:     "Test Case 1",
			fs:       os.DirFS("/tmp"),
			root:     "subdir",
			expected: os.DirFS("/tmp/subdir"),
			err:      nil,
		},
		{
			name:     "Test Case 2",
			fs:       os.DirFS("/tmp"),
			root:     "/abs/path/subdir",
			expected: os.DirFS("/abs/path/subdir"),
			err:      nil,
		},
		{
			name:     "Test Case 3",
			fs:       os.DirFS("/tmp"),
			root:     "non-existing-subdir",
			expected: nil,
			err:      &fs.PathError{Op: "open", Path: "non-existing-subdir", Err: syscall.ENOENT},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := subFS(tc.fs, tc.root)
			if err != tc.err {
				t.Errorf("Expected error %v, got %v", tc.err, err)
			}
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected result %v, got %v", tc.expected, result)
			}
		})
	}
}
