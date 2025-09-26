package commands

import (
	"errors"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/hugolib/filesystems"
	"github.com/spf13/afero"
)

func TestcopyStaticTo_1(t *testing.T) {
	testCases := []struct {
		name          string
		sourceFs      *filesystems.SourceFilesystem
		expectedCount uint64
		expectedError error
	}{
		{
			name: "Successful copy",
			sourceFs: &filesystems.SourceFilesystem{
				Name:          "static",
				Fs:            afero.NewMemMapFs(),
				SourceFs:      afero.NewOsFs(),
				PublishFolder: "public",
			},
			expectedCount: 1,
			expectedError: nil,
		},
		{
			name: "Error in copy",
			sourceFs: &filesystems.SourceFilesystem{
				Name:          "static",
				Fs:            afero.NewMemMapFs(),
				SourceFs:      afero.NewOsFs(),
				PublishFolder: "nonexistent",
			},
			expectedCount: 0,
			expectedError: errors.New("file does not exist"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			hb := &hugoBuilder{}
			count, err := hb.copyStaticTo(tc.sourceFs)

			if count != tc.expectedCount {
				t.Errorf("Expected count %d, but got %d", tc.expectedCount, count)
			}

			if (err != nil && tc.expectedError == nil) || (err == nil && tc.expectedError != nil) {
				t.Errorf("Expected error %v, but got %v", tc.expectedError, err)
			}

			if err != nil && tc.expectedError != nil && err.Error() != tc.expectedError.Error() {
				t.Errorf("Expected error message %s, but got %s", tc.expectedError.Error(), err.Error())
			}
		})
	}
}
