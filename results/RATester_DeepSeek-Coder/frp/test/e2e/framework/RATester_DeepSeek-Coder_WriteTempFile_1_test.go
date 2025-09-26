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

	fileName := "test.txt"
	content := "Hello, World!"
	filePath := f.WriteTempFile(fileName, content)

	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatal(err)
	}

	if string(data) != content {
		t.Errorf("Expected content to be %q, got %q", content, string(data))
	}
}
