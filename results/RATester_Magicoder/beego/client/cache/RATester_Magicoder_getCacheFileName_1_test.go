package cache

import (
	"fmt"
	"testing"
)

func TestgetCacheFileName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fc := &FileCache{
		CachePath:      "/tmp/cache",
		FileSuffix:     ".txt",
		DirectoryLevel: 2,
	}

	key := "test_key"
	expectedFileName := "/tmp/cache/74/15/test_key.txt"

	fileName, err := fc.getCacheFileName(key)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if fileName != expectedFileName {
		t.Errorf("Expected file name %s, but got %s", expectedFileName, fileName)
	}
}
