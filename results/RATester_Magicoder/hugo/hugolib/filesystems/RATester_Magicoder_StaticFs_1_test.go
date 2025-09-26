package filesystems

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/hugofs"
	"github.com/spf13/afero"
)

func TestStaticFs_1(t *testing.T) {
	testCases := []struct {
		name     string
		lang     string
		source   SourceFilesystems
		expected afero.Fs
	}{
		{
			name:     "Test case 1: When lang exists in Static",
			lang:     "en",
			source:   SourceFilesystems{Static: map[string]*SourceFilesystem{"en": {Fs: afero.NewMemMapFs()}}},
			expected: afero.NewMemMapFs(),
		},
		{
			name:     "Test case 2: When lang does not exist in Static but empty string exists",
			lang:     "en",
			source:   SourceFilesystems{Static: map[string]*SourceFilesystem{"": {Fs: afero.NewMemMapFs()}}},
			expected: afero.NewMemMapFs(),
		},
		{
			name:     "Test case 3: When lang and empty string does not exist in Static",
			lang:     "en",
			source:   SourceFilesystems{Static: map[string]*SourceFilesystem{}},
			expected: hugofs.NoOpFs,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.source.StaticFs(tc.lang)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
