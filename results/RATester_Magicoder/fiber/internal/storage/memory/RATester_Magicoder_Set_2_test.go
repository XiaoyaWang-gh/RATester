package memory

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestSet_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	storage := &Storage{
		db: make(map[string]entry),
	}

	key := "testKey"
	val := []byte("testValue")
	exp := time.Duration(10)

	err := storage.Set(key, val, exp)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	storage.mux.RLock()
	defer storage.mux.RUnlock()

	entry, ok := storage.db[key]
	if !ok {
		t.Errorf("Expected key %s to exist in the storage", key)
	}

	if !bytes.Equal(entry.data, val) {
		t.Errorf("Expected value %v, but got %v", val, entry.data)
	}

	if entry.expiry == 0 {
		t.Errorf("Expected expiry to be set, but got 0")
	}
}
