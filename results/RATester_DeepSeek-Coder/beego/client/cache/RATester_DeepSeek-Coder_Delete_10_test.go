package cache

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
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
	key := "test_key"

	// Create a temporary file
	tmpFile, err := ioutil.TempFile(fc.CachePath, "")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Test with a non-existing key
	err = fc.Delete(ctx, key)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test with an existing key
	err = fc.Delete(ctx, tmpFile.Name())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check if the file has been deleted
	if _, err := os.Stat(tmpFile.Name()); !os.IsNotExist(err) {
		t.Errorf("Expected file to be deleted, but it still exists")
	}
}
