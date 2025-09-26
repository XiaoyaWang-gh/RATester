package common

import (
	"fmt"
	"testing"
)

func TestFileExists_1(t *testing.T) {
	testCases := []struct {
		name     string
		fileName string
		want     bool
	}{
		{"File exists", "/path/to/existing/file", true},
		{"File does not exist", "/path/to/non-existing/file", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := FileExists(tc.fileName)
			if got != tc.want {
				t.Errorf("FileExists(%s) = %t; want %t", tc.fileName, got, tc.want)
			}
		})
	}
}
