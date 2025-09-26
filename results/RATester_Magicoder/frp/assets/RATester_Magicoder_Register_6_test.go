package assets

import (
	"fmt"
	"io"
	"testing"
	"testing/fstest"
)

func TestRegister_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testFS := fstest.MapFS{
		"static/file1.txt": {Data: []byte("content1")},
		"static/file2.txt": {Data: []byte("content2")},
	}

	Register(testFS)

	// Test if the content is correctly set
	file1, err := content.Open("file1.txt")
	if err != nil {
		t.Errorf("Failed to open file1.txt: %v", err)
	}
	defer file1.Close()

	file2, err := content.Open("file2.txt")
	if err != nil {
		t.Errorf("Failed to open file2.txt: %v", err)
	}
	defer file2.Close()

	// Test if the content is correct
	data, err := io.ReadAll(file1)
	if err != nil {
		t.Errorf("Failed to read file1.txt: %v", err)
	}
	if string(data) != "content1" {
		t.Errorf("Expected content1, got %s", string(data))
	}

	data, err = io.ReadAll(file2)
	if err != nil {
		t.Errorf("Failed to read file2.txt: %v", err)
	}
	if string(data) != "content2" {
		t.Errorf("Expected content2, got %s", string(data))
	}
}
