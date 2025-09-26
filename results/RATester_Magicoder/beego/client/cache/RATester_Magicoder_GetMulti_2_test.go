package cache

import (
	"context"
	"fmt"
	"testing"
)

func TestGetMulti_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fc := &FileCache{
		CachePath:      "/tmp/test",
		FileSuffix:     ".cache",
		DirectoryLevel: 2,
	}

	keys := []string{"key1", "key2", "key3"}
	values := []interface{}{"value1", "value2", "value3"}

	for i, key := range keys {
		err := fc.Put(context.Background(), key, values[i], 0)
		if err != nil {
			t.Fatalf("Failed to put key: %s, value: %v, error: %v", key, values[i], err)
		}
	}

	rc, err := fc.GetMulti(context.Background(), keys)
	if err != nil {
		t.Fatalf("Failed to get multi keys: %v, error: %v", keys, err)
	}

	for i, key := range keys {
		if rc[i] != values[i] {
			t.Fatalf("Failed to get correct value for key: %s, expected: %v, got: %v", key, values[i], rc[i])
		}
	}
}
