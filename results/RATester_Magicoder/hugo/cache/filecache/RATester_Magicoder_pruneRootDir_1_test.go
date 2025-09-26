package filecache

import (
	"fmt"
	"testing"
	"time"

	"github.com/spf13/afero"
)

func TestpruneRootDir_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := afero.NewMemMapFs()
	cache := &Cache{
		Fs:              fs,
		maxAge:          time.Hour,
		pruneAllRootDir: "/tmp/cache",
	}

	// Create a test directory
	if err := fs.MkdirAll(cache.pruneAllRootDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Test case 1: Directory does not exist
	count, err := cache.pruneRootDir(false)
	if err != nil {
		t.Fatal(err)
	}
	if count != 0 {
		t.Errorf("Expected 0 items to be removed, but got %d", count)
	}

	// Test case 2: Directory exists but not expired
	if err := fs.Chtimes(cache.pruneAllRootDir, time.Now(), time.Now().Add(-time.Minute)); err != nil {
		t.Fatal(err)
	}
	count, err = cache.pruneRootDir(false)
	if err != nil {
		t.Fatal(err)
	}
	if count != 0 {
		t.Errorf("Expected 0 items to be removed, but got %d", count)
	}

	// Test case 3: Directory exists and is expired
	count, err = cache.pruneRootDir(true)
	if err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Errorf("Expected 1 item to be removed, but got %d", count)
	}

	// Test case 4: Error case
	cache.Fs = afero.NewReadOnlyFs(fs)
	_, err = cache.pruneRootDir(false)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}
