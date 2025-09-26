package httplib

import (
	"fmt"
	"testing"
)

func TestPathExistAndMkdir_1(t *testing.T) {
	testCases := []struct {
		name     string
		filename string
		wantErr  bool
	}{
		{
			name:     "existing directory",
			filename: "testdata/existing",
			wantErr:  false,
		},
		{
			name:     "non-existing directory",
			filename: "testdata/non-existing",
			wantErr:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := pathExistAndMkdir(tc.filename)
			if (err != nil) != tc.wantErr {
				t.Errorf("pathExistAndMkdir() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
