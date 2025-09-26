package framework

import (
	"fmt"
	"os"
	"testing"
)

func TestWriteTempFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tempDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	f := &Framework{
		TempDirectory: tempDir,
	}

	filePath := f.WriteTempFile("test.txt", "test content")

	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatal(err)
	}

	if string(content) != "test content" {
		t.Fatalf("Expected content 'test content', got '%s'", string(content))
	}
}
