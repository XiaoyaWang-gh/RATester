package filecache

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestgetBytesAndRemoveIfExpired_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := afero.NewMemMapFs()
	cache := &Cache{
		Fs: fs,
	}

	// Create a file in the cache
	file, err := fs.Create("testfile")
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	_, err = file.Write([]byte("test data"))
	if err != nil {
		t.Fatalf("Failed to write to test file: %v", err)
	}
	file.Close()

	// Test the method
	data, removed := cache.getBytesAndRemoveIfExpired("testfile")
	if removed {
		t.Errorf("Expected file not to be removed, but it was")
	}
	if string(data) != "test data" {
		t.Errorf("Expected 'test data', but got '%s'", string(data))
	}

	// Test the method with a non-existent file
	data, removed = cache.getBytesAndRemoveIfExpired("nonexistent")
	if removed {
		t.Errorf("Expected file not to be removed, but it was")
	}
	if data != nil {
		t.Errorf("Expected nil data, but got '%s'", string(data))
	}

	// Test the method with a file that can't be read
	file, err = fs.Create("unreadable")
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	file.Close()
	data, removed = cache.getBytesAndRemoveIfExpired("unreadable")
	if removed {
		t.Errorf("Expected file not to be removed, but it was")
	}
	if data != nil {
		t.Errorf("Expected nil data, but got '%s'", string(data))
	}
}
