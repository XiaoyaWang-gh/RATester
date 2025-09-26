package deploy

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deploy/deployconfig"
	"github.com/gohugoio/hugo/media"
	"github.com/spf13/afero"
)

func TestNewLocalFile_1(t *testing.T) {
	t.Parallel()

	memFs := afero.NewMemMapFs()
	testFileContent := "This is a test file"

	testCases := []struct {
		name        string
		fileContent string
		matcher     *deployconfig.Matcher
		mediaTypes  media.Types
		expectErr   bool
	}{
		{
			name:        "valid file",
			fileContent: testFileContent,
			matcher:     nil,
			mediaTypes:  nil,
			expectErr:   false,
		},
		{
			name:        "invalid file",
			fileContent: "",
			matcher:     nil,
			mediaTypes:  nil,
			expectErr:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := afero.WriteFile(memFs, "testfile", []byte(tc.fileContent), 0644)
			if err != nil {
				t.Fatalf("Failed to create test file: %v", err)
			}

			_, err = newLocalFile(memFs, "testfile", "testfile", tc.matcher, tc.mediaTypes)
			if (err != nil) != tc.expectErr {
				t.Errorf("Expected error: %v, got: %v", tc.expectErr, err)
			}
		})
	}
}
