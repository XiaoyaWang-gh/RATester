package os

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestFileExists_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{
		readFileFs: afero.NewMemMapFs(),
		workFs:     afero.NewMemMapFs(),
	}

	// Test case 1: File exists
	err := ns.readFileFs.MkdirAll("test", 0755)
	if err != nil {
		t.Fatal(err)
	}
	_, err = ns.readFileFs.Create("test/file.txt")
	if err != nil {
		t.Fatal(err)
	}
	exists, err := ns.FileExists("test/file.txt")
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Error("Expected file to exist, but it does not")
	}

	// Test case 2: File does not exist
	exists, err = ns.FileExists("test/nonexistent.txt")
	if err != nil {
		t.Fatal(err)
	}
	if exists {
		t.Error("Expected file to not exist, but it does")
	}

	// Test case 3: Invalid path
	_, err = ns.FileExists(123)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}
