package cache

import (
	"fmt"
	"testing"
)

func TestGetCacheFileName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fc := &FileCache{
		CachePath:      "/tmp/test",
		FileSuffix:     ".txt",
		DirectoryLevel: 2,
	}

	key := "test_key"
	expectedFileName := "/tmp/test/8e/84/test_key.txt"
	fileName, err := fc.getCacheFileName(key)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if fileName != expectedFileName {
		t.Errorf("Expected %s, got %s", expectedFileName, fileName)
	}
}
