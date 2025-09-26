package logs

import (
	"fmt"
	"os"
	"testing"
)

func TestdeleteOldLog_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &fileLogWriter{
		Filename: "test.log",
		Hourly:   true,
		MaxHours: 1,
	}

	// Create a test file
	file, err := os.Create(w.Filename)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(w.Filename)

	// Write some data to the file
	_, err = file.WriteString("Test data")
	if err != nil {
		t.Fatalf("Failed to write to test file: %v", err)
	}
	file.Close()

	// Call the method under test
	w.deleteOldLog()

	// Check if the file was deleted
	_, err = os.Stat(w.Filename)
	if err == nil {
		t.Errorf("Expected file to be deleted, but it still exists")
	}
}
