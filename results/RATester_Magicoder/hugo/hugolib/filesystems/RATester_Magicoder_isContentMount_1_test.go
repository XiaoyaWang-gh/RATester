package filesystems

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/hugofs/files"
	"github.com/gohugoio/hugo/modules"
)

func TestisContentMount_1(t *testing.T) {
	b := &sourceFilesystemsBuilder{}

	testCases := []struct {
		name     string
		mnt      modules.Mount
		expected bool
	}{
		{
			name: "Content mount",
			mnt: modules.Mount{
				Target: files.ComponentFolderContent,
			},
			expected: true,
		},
		{
			name: "Non-content mount",
			mnt: modules.Mount{
				Target: "some/other/path",
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := b.isContentMount(tc.mnt)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
