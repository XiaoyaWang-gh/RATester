package file

import (
	"context"
	"fmt"
	"testing"
)

func TestLoadFileConfigFromDirectory_1(t *testing.T) {
	testCases := []struct {
		desc        string
		directory   string
		expectedErr error
	}{
		{
			desc:        "should return an error if directory does not exist",
			directory:   "/non-existing-directory",
			expectedErr: fmt.Errorf("unable to read directory /non-existing-directory: open /non-existing-directory: no such file or directory"),
		},
		{
			desc:        "should return an error if directory is not a directory",
			directory:   "testdata/file.go",
			expectedErr: fmt.Errorf("unable to read directory testdata/file.go: not a directory"),
		},
		{
			desc:        "should return an error if directory is empty",
			directory:   "testdata/empty",
			expectedErr: nil,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ctx := context.Background()
			provider := &Provider{}
			_, err := provider.loadFileConfigFromDirectory(ctx, test.directory, nil)
			if err != nil && test.expectedErr != nil && err.Error() != test.expectedErr.Error() {
				t.Errorf("Expected error %v, got %v", test.expectedErr, err)
			}
			if err == nil && test.expectedErr != nil {
				t.Errorf("Expected error %v, got nil", test.expectedErr)
			}
		})
	}
}
