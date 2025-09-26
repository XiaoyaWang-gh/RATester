package commands

import (
	"fmt"
	"os"
	"testing"
)

func TestMkdir_2(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected string
	}{
		{"Test case 1", []string{"testdata", "subdir"}, "testdata/subdir"},
		{"Test case 2", []string{"testdata", "subdir1", "subdir2"}, "testdata/subdir1/subdir2"},
		{"Test case 3", []string{"testdata", "subdir1", "subdir2", "subdir3"}, "testdata/subdir1/subdir2/subdir3"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			mkdir(tc.input...)
			_, err := os.Stat(tc.expected)
			if os.IsNotExist(err) {
				t.Errorf("Expected directory %s does not exist", tc.expected)
			}
		})
	}
}
