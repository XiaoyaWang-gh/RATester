package cache

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestPut_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fc := &FileCache{
		CachePath:  "/tmp/test",
		FileSuffix: ".cache",
	}

	ctx := context.Background()
	key := "testKey"
	val := "testValue"
	timeout := 1 * time.Hour

	err := fc.Put(ctx, key, val, timeout)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	_, err = os.Stat(filepath.Join(fc.CachePath, key+fc.FileSuffix))
	if os.IsNotExist(err) {
		t.Errorf("Expected file to exist, got %v", err)
	}
}
