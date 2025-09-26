package logs

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestDestroy_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &fileLogWriter{
		Filename: "test.log",
	}

	// Create a temporary file for testing
	tmpFile, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	w.fileWriter = tmpFile

	// Test Destroy
	w.Destroy()

	// Check if the file is closed
	_, err = w.fileWriter.Stat()
	if err == nil {
		t.Errorf("Expected error when trying to stat closed file, got nil")
	}
}
