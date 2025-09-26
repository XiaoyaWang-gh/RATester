package modules

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestrewriteGoModRewrite_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{
		fs: afero.NewMemMapFs(),
	}

	// Create a test file
	testFile := "test.txt"
	afero.WriteFile(client.fs, testFile, []byte("test line\n"), 0644)

	// Test case 1: file does not exist
	_, err := client.rewriteGoModRewrite("nonexistent.txt", map[string]bool{"test": true})
	if err == nil {
		t.Error("Expected error, got nil")
	}

	// Test case 2: file exists, no match
	_, err = client.rewriteGoModRewrite(testFile, map[string]bool{"nonexistent": true})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test case 3: file exists, match
	_, err = client.rewriteGoModRewrite(testFile, map[string]bool{"test": true})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
