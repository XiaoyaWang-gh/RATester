package echo

import (
	"fmt"
	"io/fs"
	"os"
	"reflect"
	"testing"
)

func TestMustSubFS_1(t *testing.T) {
	testCases := []struct {
		name     string
		fs       fs.FS
		fsRoot   string
		expected fs.FS
	}{
		{
			name:     "valid fs and root",
			fs:       os.DirFS("/tmp"),
			fsRoot:   "test",
			expected: os.DirFS("/tmp/test"),
		},
		{
			name:     "invalid fs",
			fs:       os.DirFS("/non-existing-dir"),
			fsRoot:   "test",
			expected: nil,
		},
		{
			name:     "invalid root",
			fs:       os.DirFS("/tmp"),
			fsRoot:   "non-existing-dir",
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			defer func() {
				if r := recover(); r != nil {
					if tc.expected == nil {
						return
					}
					t.Errorf("expected panic, but got: %v", r)
				}
			}()

			result := MustSubFS(tc.fs, tc.fsRoot)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
