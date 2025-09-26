package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func Testmkdir_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testDir := "testDir"
	testSubDir := "testSubDir"
	testPath := filepath.Join(testDir, testSubDir)

	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Change the current working directory to the temporary directory
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	// Call the function under test
	mkdir(testDir, testSubDir)

	// Check if the directory was created
	if _, err := os.Stat(testPath); os.IsNotExist(err) {
		t.Errorf("Expected directory %s does not exist", testPath)
	}
}
