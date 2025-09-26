package httplib

import (
	"fmt"
	"os"
	"testing"
)

func TestToFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BeegoHTTPRequest{
		url: "http://example.com",
	}

	filename := "test.txt"
	err := b.ToFile(filename)
	if err != nil {
		t.Errorf("ToFile failed: %v", err)
	}

	// Check if the file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Errorf("File %s does not exist", filename)
	}

	// Clean up
	err = os.Remove(filename)
	if err != nil {
		t.Errorf("Failed to remove file %s: %v", filename, err)
	}
}
