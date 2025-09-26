package cache

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestDelete_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fc := &FileCache{
		CachePath: "/tmp/test",
	}

	ctx := context.Background()

	// Create a test file
	testKey := "testKey"
	testFile := filepath.Join(fc.CachePath, testKey)
	err := os.WriteFile(testFile, []byte("test data"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Test deleting the file
	err = fc.Delete(ctx, testKey)
	if err != nil {
		t.Errorf("Delete failed: %v", err)
	}

	// Check if the file still exists
	_, err = os.Stat(testFile)
	if err == nil {
		t.Errorf("File still exists after delete")
	}
}
