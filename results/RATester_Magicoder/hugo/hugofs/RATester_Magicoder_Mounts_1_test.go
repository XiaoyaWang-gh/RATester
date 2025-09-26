package hugofs

import (
	"fmt"
	"testing"

	radix "github.com/armon/go-radix"
	"github.com/spf13/afero"
)

func TestMounts_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := &RootMappingFs{
		id:            "test",
		Fs:            afero.NewMemMapFs(),
		rootMapToReal: radix.New(),
		realMapToRoot: radix.New(),
	}

	// Create some test files
	err := fs.Fs.MkdirAll("/test/dir1", 0755)
	if err != nil {
		t.Fatal(err)
	}
	err = afero.WriteFile(fs.Fs, "/test/dir1/file1", []byte("test1"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	err = fs.Fs.MkdirAll("/test/dir2", 0755)
	if err != nil {
		t.Fatal(err)
	}
	err = afero.WriteFile(fs.Fs, "/test/dir2/file2", []byte("test2"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Test Mounts function
	files, err := fs.Mounts("/test")
	if err != nil {
		t.Fatal(err)
	}

	// Check if the correct number of files is returned
	if len(files) != 2 {
		t.Fatalf("Expected 2 files, got %d", len(files))
	}

	// Check if the correct files are returned
	if files[0].Name() != "dir1" || files[1].Name() != "dir2" {
		t.Fatalf("Expected files 'dir1' and 'dir2', got %s and %s", files[0].Name(), files[1].Name())
	}
}
