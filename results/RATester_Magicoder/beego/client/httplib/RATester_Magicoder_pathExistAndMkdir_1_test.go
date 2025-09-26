package httplib

import (
	"fmt"
	"testing"
)

func TestpathExistAndMkdir_1(t *testing.T) {
	testCases := []struct {
		name     string
		filename string
		wantErr  bool
	}{
		{
			name:     "Existing Directory",
			filename: "existing_dir",
			wantErr:  false,
		},
		{
			name:     "Non-Existing Directory",
			filename: "non_existing_dir",
			wantErr:  false,
		},
		{
			name:     "Invalid Directory",
			filename: "invalid_dir",
			wantErr:  true,
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
				return
			}
		})
	}
}
