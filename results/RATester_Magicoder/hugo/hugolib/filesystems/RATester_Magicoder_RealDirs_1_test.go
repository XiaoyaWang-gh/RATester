package filesystems

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestRealDirs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := &SourceFilesystem{
		Name: "test",
		Fs:   afero.NewMemMapFs(),
	}

	// Create some test directories
	fs.Fs.MkdirAll("/test/dir1", 0755)
	fs.Fs.MkdirAll("/test/dir2", 0755)
	fs.Fs.MkdirAll("/test/dir3", 0755)

	// Test the RealDirs function
	dirs := fs.RealDirs("dir")
	if len(dirs) != 3 {
		t.Errorf("Expected 3 directories, got %d", len(dirs))
	}

	// Test with a non-existent directory
	dirs = fs.RealDirs("nonexistent")
	if len(dirs) != 0 {
		t.Errorf("Expected 0 directories, got %d", len(dirs))
	}
}
