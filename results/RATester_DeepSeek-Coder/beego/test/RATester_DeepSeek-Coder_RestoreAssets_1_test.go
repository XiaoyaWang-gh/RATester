package test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestRestoreAssets_1(t *testing.T) {
	dir, err := ioutil.TempDir("", "TestRestoreAssets")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	testCases := []struct {
		name     string
		expected error
	}{
		{
			name:     "existing_directory",
			expected: nil,
		},
		{
			name:     "non_existing_directory",
			expected: errors.New("some error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := RestoreAssets(dir, tc.name)
			if (err != nil) != (tc.expected != nil) {
				t.Errorf("expected error %v, got %v", tc.expected, err)
			}
		})
	}
}
