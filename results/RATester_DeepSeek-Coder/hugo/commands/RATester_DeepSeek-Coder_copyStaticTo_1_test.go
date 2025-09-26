package commands

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/hugolib/filesystems"
	"github.com/spf13/afero"
)

func TestCopyStaticTo_1(t *testing.T) {
	testCases := []struct {
		name          string
		sourceFs      *filesystems.SourceFilesystem
		expectedError error
	}{
		{
			name: "Successful copy",
			sourceFs: &filesystems.SourceFilesystem{
				Name:          "static",
				Fs:            afero.NewMemMapFs(),
				SourceFs:      afero.NewOsFs(),
				PublishFolder: "",
			},
			expectedError: nil,
		},
		{
			name: "Copy with PublishFolder",
			sourceFs: &filesystems.SourceFilesystem{
				Name:          "static",
				Fs:            afero.NewMemMapFs(),
				SourceFs:      afero.NewOsFs(),
				PublishFolder: "publish",
			},
			expectedError: nil,
		},
		{
			name: "Copy with non-existing sourceFs",
			sourceFs: &filesystems.SourceFilesystem{
				Name:          "non-existing",
				Fs:            afero.NewMemMapFs(),
				SourceFs:      afero.NewOsFs(),
				PublishFolder: "",
			},
			expectedError: fmt.Errorf("source filesystem not found: non-existing"),
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
			_, err := hb.copyStaticTo(tc.sourceFs)
			if err != tc.expectedError {
				t.Errorf("Expected error %v, got %v", tc.expectedError, err)
			}
		})
	}
}
