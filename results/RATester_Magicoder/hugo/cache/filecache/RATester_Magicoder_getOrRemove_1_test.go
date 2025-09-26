package filecache

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestgetOrRemove_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := afero.NewMemMapFs()
	cache := &Cache{
		Fs: fs,
	}

	testFile := "test.txt"
	testContent := "Hello, World!"

	// Create a test file
	err := afero.WriteFile(fs, testFile, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Test getOrRemove with a valid file
	file := cache.getOrRemove(testFile)
	if file == nil {
		t.Fatalf("Expected a valid file, but got nil")
	}

	content, err := afero.ReadFile(fs, testFile)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	if string(content) != testContent {
		t.Fatalf("Expected content '%s', but got '%s'", testContent, string(content))
	}

	// Test getOrRemove with an invalid file
	invalidFile := cache.getOrRemove("invalid.txt")
	if invalidFile != nil {
		t.Fatalf("Expected nil, but got a valid file")
	}
}
