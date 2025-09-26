package logs

import (
	"fmt"
	"os"
	"testing"
)

func TestLines_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &fileLogWriter{
		Filename: "test.log",
	}

	// Create a test file
	fd, err := os.Create(w.Filename)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(w.Filename)
	defer fd.Close()

	// Write some test data to the file
	testData := []byte("line1\nline2\nline3\n")
	_, err = fd.Write(testData)
	if err != nil {
		t.Fatalf("Failed to write test data to file: %v", err)
	}

	// Test the lines function
	count, err := w.lines()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if count != 3 {
		t.Errorf("Expected 3 lines, got %d", count)
	}
}
