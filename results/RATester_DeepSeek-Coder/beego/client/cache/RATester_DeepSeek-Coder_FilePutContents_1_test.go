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

	testFile := "test.txt"
	testContent := []byte("Hello, World!\n")

	err := FilePutContents(testFile, testContent)
	if err != nil {
		t.Errorf("FilePutContents() error = %v, wantErr nil", err)
		return
	}

	data, err := os.ReadFile(testFile)
	if err != nil {
		t.Errorf("Failed to read test file: %v", err)
		return
	}

	if string(data) != string(testContent) {
		t.Errorf("FilePutContents() = %v, want %v", string(data), string(testContent))
	}

	os.Remove(testFile)
}
