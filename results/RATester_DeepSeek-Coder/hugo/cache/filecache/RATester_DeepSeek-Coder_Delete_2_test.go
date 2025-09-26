package filecache

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestDelete_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := &Cache{
		Fs: afero.NewMemMapFs(),
	}
	httpC := &httpCache{
		c: cache,
	}

	key := "testKey"

	// Create a file with the key
	err := afero.WriteFile(cache.Fs, key, []byte("test data"), 0644)
	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	// Check that the file exists before deletion
	_, err = cache.Fs.Stat(key)
	if err != nil {
		t.Fatalf("File does not exist before deletion: %v", err)
	}

	// Delete the file
	httpC.Delete(key)

	// Check that the file does not exist after deletion
	_, err = cache.Fs.Stat(key)
	if err == nil {
		t.Fatalf("File exists after deletion")
	}
}
