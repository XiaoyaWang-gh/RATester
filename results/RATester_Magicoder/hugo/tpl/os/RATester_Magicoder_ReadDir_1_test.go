package os

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
	"github.com/spf13/afero"
)

func TestReadDir_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{
		readFileFs: afero.NewMemMapFs(),
		workFs:     afero.NewMemMapFs(),
		deps:       &deps.Deps{},
	}

	// Create a directory
	err := ns.workFs.Mkdir("/test", 0755)
	if err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}

	// Create a file in the directory
	err = afero.WriteFile(ns.workFs, "/test/file1.txt", []byte("content1"), 0644)
	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	// Test the ReadDir function
	files, err := ns.ReadDir("/test")
	if err != nil {
		t.Fatalf("ReadDir failed: %v", err)
	}

	if len(files) != 1 {
		t.Fatalf("Expected 1 file, got %d", len(files))
	}

	if files[0].Name() != "file1.txt" {
		t.Fatalf("Expected file name 'file1.txt', got '%s'", files[0].Name())
	}
}
