package acme

import (
	"fmt"
	"testing"
)

func TestNewLocalStore_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	filename := "test.db"
	store := NewLocalStore(filename)

	if store.filename != filename {
		t.Errorf("Expected filename to be %s, but got %s", filename, store.filename)
	}

	if store.saveDataChan == nil {
		t.Error("Expected saveDataChan to be initialized")
	}

	if store.storedData == nil {
		t.Error("Expected storedData to be initialized")
	}
}
