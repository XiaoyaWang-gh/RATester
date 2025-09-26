package install

import (
	"fmt"
	"testing"
)

func TestCopyDir_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	srcPath := "/path/to/source/directory"
	destPath := "/path/to/destination/directory"

	// Test case 1: Source directory does not exist
	err := CopyDir(srcPath, destPath)
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	// Test case 2: Destination directory does not exist
	srcPath = "/path/to/existing/directory"
	err = CopyDir(srcPath, destPath)
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	// Test case 3: Source and destination are the same
	srcPath = "/path/to/existing/directory"
	destPath = "/path/to/existing/directory"
	err = CopyDir(srcPath, destPath)
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	// Test case 4: Source and destination are different and valid
	srcPath = "/path/to/existing/directory"
	destPath = "/path/to/new/directory"
	err = CopyDir(srcPath, destPath)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
