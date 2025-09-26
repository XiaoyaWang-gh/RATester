package filecache

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestremoveIfExpired_1(t *testing.T) {
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
	fileName := "test.txt"
	file, _ := fs.Create(fileName)
	file.WriteString("test content")
	file.Close()

	// Test the method
	removed, err := cache.removeIfExpired(fileName)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if removed {
		t.Errorf("Expected file not to be removed")
	}

	// Set maxAge to a negative value to make the file expired
	cache.maxAge = -1
	removed, err = cache.removeIfExpired(fileName)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !removed {
		t.Errorf("Expected file to be removed")
	}

	// Check if the file is actually removed
	_, err = fs.Stat(fileName)
	if err == nil {
		t.Errorf("Expected file to be removed")
	}
}
