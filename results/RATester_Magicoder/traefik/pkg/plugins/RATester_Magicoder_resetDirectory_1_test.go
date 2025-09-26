package plugins

import (
	"fmt"
	"os"
	"testing"
)

func TestResetDirectory_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	dir := "test_dir"
	defer os.RemoveAll(dir)

	// Test case 1: Successful reset
	err := resetDirectory(dir)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Test case 2: Error getting absolute path
	dir = "~"
	err = resetDirectory(dir)
	if err == nil {
		t.Error("Expected an error, but got none")
	}

	// Test case 3: Error getting current directory
	dir = "."
	err = resetDirectory(dir)
	if err == nil {
		t.Error("Expected an error, but got none")
	}

	// Test case 4: Error when directory is parent of current path
	dir = ".."
	err = resetDirectory(dir)
	if err == nil {
		t.Error("Expected an error, but got none")
	}

	// Test case 5: Error when removing directory
	dir = "/"
	err = resetDirectory(dir)
	if err == nil {
		t.Error("Expected an error, but got none")
	}

	// Test case 6: Error when creating directory
	dir = "/dev/null"
	err = resetDirectory(dir)
	if err == nil {
		t.Error("Expected an error, but got none")
	}
}
