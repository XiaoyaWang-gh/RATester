package cache

import (
	"fmt"
	"os"
	"testing"
)

func TestFilePutContents_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	filename := "test.txt"
	content := []byte("Hello, World!")

	err := FilePutContents(filename, content)
	if err != nil {
		t.Errorf("FilePutContents failed: %v", err)
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		t.Errorf("Failed to read file: %v", err)
	}

	if string(data) != string(content) {
		t.Errorf("File content does not match: expected %s, got %s", string(content), string(data))
	}

	err = os.Remove(filename)
	if err != nil {
		t.Errorf("Failed to remove file: %v", err)
	}
}
