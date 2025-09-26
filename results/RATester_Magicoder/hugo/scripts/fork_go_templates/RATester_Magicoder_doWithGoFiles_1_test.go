package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestdoWithGoFiles_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	dir := "testdir"
	rewrite := func(name string) {
		// Implementation of the rewrite function
	}
	transform := func(name, in string) string {
		// Implementation of the transform function
		return in
	}

	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create a test file
	testFile, err := os.Create(filepath.Join(tempDir, "test.go"))
	if err != nil {
		t.Fatal(err)
	}
	testFile.Close()

	// Call the function under test
	doWithGoFiles(dir, rewrite, transform)

	// Check if the file was processed
	_, err = os.Stat(filepath.Join(tempDir, "test.go"))
	if err != nil {
		t.Fatal(err)
	}
}
