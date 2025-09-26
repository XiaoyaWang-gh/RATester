package filecache

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestGetString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := afero.NewMemMapFs()
	cache := &Cache{
		Fs: fs,
	}

	id := "test_id"
	content := "test_content"

	// Create a file with the content
	err := afero.WriteFile(fs, id, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	// Call the method under test
	result := cache.GetString(id)

	// Verify the result
	if result != content {
		t.Errorf("Expected %s, got %s", content, result)
	}
}
