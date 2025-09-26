package filecache

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestPrune_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := afero.NewMemMapFs()
	cache := &Cache{
		Fs: fs,
	}

	// Create some files in the cache
	for i := 0; i < 10; i++ {
		afero.WriteFile(fs, fmt.Sprintf("file%d", i), []byte("content"), 0644)
	}

	// Call the method under test
	counter, err := cache.Prune(false)

	// Check the result
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if counter != 0 {
		t.Errorf("Expected 0 files to be removed, but %d were", counter)
	}

	// Check that the files still exist
	for i := 0; i < 10; i++ {
		_, err := fs.Stat(fmt.Sprintf("file%d", i))
		if err != nil {
			t.Errorf("File file%d was removed", i)
		}
	}
}
